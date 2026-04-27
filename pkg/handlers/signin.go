package handlers

import (
	"encoding/json"
	"go-todo-project/pkg/auth"
	"net/http"
	"os"
)

type SigninRequest struct {
	Password string `json:"password"`
}

// SigninHandler обрабатывает POST запрос пароля /api/signin
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	expectPass := os.Getenv("TODO_PASSWORD")
	if expectPass == "" {
		writeJSON(w, map[string]string{"token": "no-auth"})
		return
	}

	var request SigninRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(w, map[string]string{"error": "incorect JSON"})
		return
	}

	if request.Password != expectPass {
		writeJSON(w, map[string]string{"error": "wrong password"})
		return
	}
	token, err := auth.GenerateToken(expectPass)
	if err != nil {
		writeJSON(w, map[string]string{"error": "token gen failed"})
		return
	}

	writeJSON(w, map[string]string{"token": token})
}
