package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandleTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		GetAllTodos(w, r)
	} else if r.Method == http.MethodPost {
		CreateNewTask(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTodoById(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/todos/")

	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		GetTodoById(w, r, id)
	} else if r.Method == http.MethodPatch {
		UpdateTodoById(w, r, id)
	} else if r.Method == http.MethodDelete {
		DeleteTodoById(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskItems, err := ReadJsonFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(taskItems, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

func GetTodoById(w http.ResponseWriter, r *http.Request, taskId uint64) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskItems, err := ReadJsonFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundtask *Task
	for _, task := range taskItems {
		if uint64(task.Id) == taskId {
			foundtask = &task
			break
		}
	}

	if foundtask == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	jsonData, err := json.MarshalIndent(foundtask, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonData))
}

func UpdateTodoById(w http.ResponseWriter, r *http.Request, taskId uint64) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskItems, err := ReadJsonFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundId int = -1
	for id, task := range taskItems {
		if uint64(task.Id) == taskId {
			foundId = id
			break
		}
	}

	if foundId == -1 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updatePayload Task

	err = json.Unmarshal(body, &updatePayload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskItems[foundId].Description = updatePayload.Description
	taskItems[foundId].Done = updatePayload.Done
	taskItems[foundId].UpdatedAt = time.Now()

	err = WriteJsonFile(taskItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(taskItems[foundId], "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/JSON")
	w.Write(jsonData)
}

func CreateNewTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPartialContent)
		return
	}

	var newTask Task
	err = json.Unmarshal(body, &newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskItems, err := ReadJsonFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskItems = append(taskItems, CreateTask(newTask.Description))

	err = WriteJsonFile(taskItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.MarshalIndent(taskItems[len(taskItems)-1], "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/JSON")
	w.Write(jsonData)
}

func DeleteTodoById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/api/todos/")
	id, err := strconv.ParseUint(path, 10, 32)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	taskItems, err := ReadJsonFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundId int = -1
	for i, task := range taskItems {
		if uint64(task.Id) == id {
			foundId = i
			break
		}
	}

	if foundId == -1 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	taskItems = append(taskItems[:foundId], taskItems[foundId+1:]...)
	err = WriteJsonFile(taskItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]string{"success": "Task deleted successfully"}
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "Application/JSON")
	w.Write(jsonData)
}
