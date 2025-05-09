package main

import (
	"html/template"
	"net/http"
)

func TodoPage(w http.ResponseWriter, r *http.Request) {
	err := CreateFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	taskItems, err := ReadJsonFile()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Tasks []Task
	}{
		Tasks: taskItems,
	}

	funcMap := template.FuncMap{
		"formatTime": formatDate,
		"add":        add,
		"time":       getTime,
	}

	tmpl, err := template.New("todoTemplate").Funcs(funcMap).ParseFiles(
		"templates/index.html", "templates/todos.html", "templates/modal.html", "templates/header.html",
	)
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
	}

	tmpl = tmpl.Lookup("index.html")
	if tmpl == nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(nil, err.Error(), http.StatusInternalServerError)
	}
}
