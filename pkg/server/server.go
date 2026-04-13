package server

import (
	"fmt"
	"net/http"
)

func StartServer(port string) error {
	webDirect := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDirect)))

	address := ":" + port
	// Пометка, чтоб видеть запуск сервера
	fmt.Println("Server started on http://localhost" + address)

	return http.ListenAndServe(address, nil)
}
