package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	webDirect := "./web"
	http.Handle("/", http.FileServer(http.Dir(webDirect)))

	// Пометка, чтоб видеть запуск сервера
	fmt.Println("Server started on port :7540")

	err := http.ListenAndServe(":7540", nil)
	if err != nil {
		log.Fatal("server error: ", err)
	}
}
