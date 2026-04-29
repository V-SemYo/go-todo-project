package server

import (
	"fmt"
	"go-todo-project/pkg/api"
	"go-todo-project/pkg/handlers"
	"net/http"
	"os"
)

func StartServer(port string) error {
	api.Init()

	webDirect := "./web"
	fileServer := http.FileServer(http.Dir(webDirect))
	if os.Getenv("TODO_PASSWORD") != "" {
		http.Handle("/", handlers.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
			fileServer.ServeHTTP(w, r)
		}))
	} else {
		http.Handle("/", fileServer)
	}

	address := ":" + port
	// Пометка, чтоб видеть запуск сервера
	fmt.Println("Server started on http://localhost" + address)

	return http.ListenAndServe(address, nil)
}
