package api

import (
	"encoding/json"
	"fmt"
	"go-todo-project/pkg/db"
	"net/http"
	"time"
)

// validateDate проверяет и корректирует дату задачи по правилам ТЗ
func validateDate(task *db.Task, now time.Time) error {
	if task.Date == "" {
		task.Date = now.Format(dateFormat)
	}
	tDate, err := time.Parse(dateFormat, task.Date)
	if err != nil {
		return fmt.Errorf("wrong date format: %w", err)
	}
	if task.Repeat != "" {
		if tDate.Before(now) || tDate.Format(dateFormat) == now.Format(dateFormat) {
			if task.Repeat == "d 1" && tDate.Format(dateFormat) == now.Format(dateFormat) {
				task.Date = now.Format(dateFormat)
				return nil
			}
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return fmt.Errorf("wrong repeat rule: %w", err)
			}
			task.Date = next
		}
		return nil
	}
	if tDate.Before(now) {
		task.Date = now.Format(dateFormat)
	}
	return nil
}

// addTaskHandle обрабатывает POST-запросы на /api/task
func addTaskHandle(w http.ResponseWriter, r *http.Request) {
	var task db.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJSON(w, map[string]string{"error": "incorect JSON"})
		return
	}

	if task.Title == "" {
		writeJSON(w, map[string]string{"error": "title is empty"})
		return
	}

	now := time.Now()
	if err := validateDate(&task, now); err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJSON(w, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(w, map[string]any{"id": id})
}
