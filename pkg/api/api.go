package api

import (
	"go-todo-project/pkg/handlers"
	"net/http"
)

// Init регистрирует все API обработчики
func Init() {
	http.HandleFunc("/api/nextdate", handlers.NextDateHandler)
	http.HandleFunc("/api/task", handlers.TaskHandler)
	http.HandleFunc("/api/tasks", handlers.TasksHandler)
	http.HandleFunc("/api/task/done", handlers.DoneTaskHandler)
}
