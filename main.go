package main

import (
	"log"
	"os"

	"go-todo-project/pkg/db"
	"go-todo-project/pkg/server"
)

func main() {
	dbFile := "scheduler.db"
	if envFile := os.Getenv("TODO_DBFILE"); envFile != "" {
		dbFile = envFile
	}

	if err := db.InitDB(dbFile); err != nil {
		log.Fatal("database initialisation error ", err)
	}
	defer db.DB.Close()

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	err := server.StartServer(port)
	if err != nil {
		log.Fatal("server error: ", err)
	}
}
