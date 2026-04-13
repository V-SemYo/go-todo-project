package main

import (
	"log"
	"os"

	"go-todo-project/pkg/server"
)

func main() {
	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	err := server.StartServer(port)
	if err != nil {
		log.Fatal("server error: ", err)
	}
}
