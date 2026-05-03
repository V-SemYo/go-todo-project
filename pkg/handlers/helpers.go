package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// writeJSON отправляет данные в формате JSON с правильным заголовком Content-Type и HTTP статус-кодом
func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON encode error: %v", err)
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
	}
}
