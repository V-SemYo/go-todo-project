package handlers

import (
	"go-todo-project/pkg/db"
	"net/http"
)

// DeleteTaskHandler обрабатывает DELETE /api/task?id=<№id>
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id not specified"})
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{})
}
