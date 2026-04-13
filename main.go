package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	webDirect := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDirect)))

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	address := ":" + port

	// Пометка, чтоб видеть запуск сервера
	fmt.Println("Server started on http://localhost" + address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal("server error: ", err)
	}
}
