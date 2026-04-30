package handlers

import (
	"go-todo-project/pkg/db"
	"net/http"
)

type TaskResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// TasksHandler обрабатывает GET-запросы /api/tasks, возвращает список задач в формате JSON
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	search := r.FormValue("search")
	limit := 40

	tasks, err := db.Tasks(limit, search)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJSON(w, http.StatusOK, TaskResp{Tasks: tasks})
}
