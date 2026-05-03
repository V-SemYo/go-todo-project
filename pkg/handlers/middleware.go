package handlers

import (
	"go-todo-project/pkg/auth"
	"net/http"
	"os"
)

// AuthMiddle проверяет, что пользователь аутентифицирован, прежде чем вызвать реальный обработчик
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		expectPass := os.Getenv("TODO_PASSWORD")
		if expectPass == "" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "need authentication", http.StatusUnauthorized)
			return
		}

		validate, err := auth.ValidateToken(cookie.Value, expectPass)
		if err != nil || !validate {
			http.Error(w, "need authentication", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
