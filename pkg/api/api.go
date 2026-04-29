package api

import (
	"go-todo-project/pkg/handlers"
	"net/http"
)

// Init регистрирует все API обработчики
func Init() {
	http.HandleFunc("/api/signin", handlers.SigninHandler)

	http.HandleFunc("/api/nextdate", handlers.AuthMiddleware(handlers.NextDateHandler))
	http.HandleFunc("/api/task", handlers.AuthMiddleware(handlers.TaskHandler))
	http.HandleFunc("/api/tasks", handlers.AuthMiddleware(handlers.TasksHandler))
	http.HandleFunc("/api/task/done", handlers.AuthMiddleware(handlers.DoneTaskHandler))
}
