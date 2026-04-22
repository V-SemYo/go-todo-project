package handlers

import (
	"go-todo-project/pkg/db"
	"net/http"
)

// GetTaskHandler обрабатывает GET /api/task?id=номер задачи
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "id not specified"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, task)
}
