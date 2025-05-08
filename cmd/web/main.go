package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/api/todos", HandleTodos)
	http.HandleFunc("/api/todos/{id}", HandleTodoById)

	// handle unknown routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		TodoPage(w, r)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
