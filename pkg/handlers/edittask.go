package handlers

import (
	"encoding/json"
	"go-todo-project/pkg/db"
	"net/http"
	"time"
)

// EditTaskHandler обрабатывает PUT /api/task
func EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "incorect JSON"})
		return
	}

	if task.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "title is empty"})
		return
	}

	now := time.Now()
	if err := validateDate(&task, now); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	if err := db.UpdateTask(&task); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}
