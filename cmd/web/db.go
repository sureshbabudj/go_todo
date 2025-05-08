package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uint32    `json:"id"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var dirPath = "data"
var filePath = filepath.Join(dirPath, "todos.json")

func CreateFile() error {
	var taskItems = []Task{CreateTask("Watch Go Crash Course"), CreateTask("Watch Go Full Course"), CreateTask("Reward myself with donut")}

	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirPath, 0755)
		if err != nil {
			return err
		}
		CreateJsonFile(taskItems)
	} else {
		_, err = os.Stat(filePath)
		if os.IsNotExist(err) {
			CreateJsonFile(taskItems)
		}
	}

	return nil
}

func CreateJsonFile(taskItems []Task) error {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println(("Error in creating json file"))
		return err
	}
	defer file.Close()

	WriteJsonFile(taskItems)

	return nil
}

func ReadJsonFile() ([]Task, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var taskItems []Task

	err = json.Unmarshal(jsonData, &taskItems)
	if err != nil {
		return nil, err
	}

	return taskItems, nil
}

func WriteJsonFile(taskItems []Task) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.MarshalIndent(taskItems, "", "   ")
	if err != nil {
		return err
	}
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func CreateTask(taskItem string) Task {
	uuidInstance, _ := uuid.NewRandom()
	id := uuidInstance.ID()
	return Task{
		Id:          id,
		Description: taskItem,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
