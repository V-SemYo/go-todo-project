package api

import (
	"net/http"
)

// taskHandler - главный обработчик для /api/task, вызывает обработчик в зависимости от HTTP-метода
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		addTaskHandle(w, r)
	default:
		writeJSON(w, map[string]string{"error": "method not allowed"})
	}
}
