package server

import (
	"fmt"
	"go-todo-project/pkg/api"
	"net/http"
)

func StartServer(port string) error {
	api.Init()

	http.Handle("/", http.FileServer(http.Dir("./web")))

	address := ":" + port
	// Пометка, чтоб видеть запуск сервера
	fmt.Println("Server started on http://localhost" + address)
	return http.ListenAndServe(address, nil)
}
