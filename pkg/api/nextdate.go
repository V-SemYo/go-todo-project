package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("parsing date error: %w", err)
	}
	if repeat == "" {
		return "", fmt.Errorf("repeat is empty")
	}
	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {
	case "d":
		if len(parts) < 2 {
			return "", fmt.Errorf("missing days number for 'd' rule")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("wrong days number: %w", err)
		}
		if days < 1 || days > 400 {
			return "", fmt.Errorf("days number must be 1 to 400")
		}
		date = date.AddDate(0, 0, days)
		for !date.After(now) {
			date = date.AddDate(0, 0, days)
		}
		return date.Format(dateFormat), nil

	case "y":
		date = date.AddDate(1, 0, 0)
		for !date.After(now) {
			date = date.AddDate(1, 0, 0)
		}
		return date.Format(dateFormat), nil
	default:
		return "", fmt.Errorf("wrong repeat rule: %s", rule)
	}
}

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	date := r.FormValue("date")
	repeat := r.FormValue("repeat")
	nowStr := r.FormValue("now")

	var now time.Time
	var err error

	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(w, "wrong 'now' format", http.StatusBadRequest)
			return
		}
	}

	next, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(next))
}
