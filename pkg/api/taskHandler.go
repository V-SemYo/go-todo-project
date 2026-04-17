package api

import (
	"net/http"
)

// TaskHandler - главный обработчик для /api/task, вызывает обработчик в зависимости от HTTP-метода
func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)
	default:
		writeJSON(w, map[string]string{"error": "method not allowed"})
	}
}
