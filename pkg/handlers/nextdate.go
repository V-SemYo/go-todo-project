package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

// NextDate вычисляет следующую дату выполнения задачи на основе правила повторения
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
		return HandlerRuleD(now, date, parts)
	case "y":
		return HandlerRuleY(now, date)
	case "w":
		return HandlerRuleW(now, date, parts)
	case "m":
		return HandlerRuleM(now, date, parts)
	default:
		return "", fmt.Errorf("wrong repeat rule: %s", rule)
	}
}

// HandlerRuleD обработчик для правила "d" (дни)
func HandlerRuleD(now, date time.Time, parts []string) (string, error) {
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
}

// HandlerRuleY обработчик для правила "y" (годы)
func HandlerRuleY(now, date time.Time) (string, error) {
	date = date.AddDate(1, 0, 0)
	for !date.After(now) {
		date = date.AddDate(1, 0, 0)
	}
	return date.Format(dateFormat), nil
}

// HandlerRuleW обработчик для правила "w" (дни недели)
func HandlerRuleW(now, date time.Time, parts []string) (string, error) {
	if len(parts) < 2 {
		return "", fmt.Errorf("missing days number for 'w' rule")
	}
	weekdays, err := ParseWeekdays(parts[1])
	if err != nil {
		return "", err
	}
	date = date.AddDate(0, 0, 1)
	for {
		if weekdays[date.Weekday()] && date.After(now) {
			return date.Format(dateFormat), nil
		}
		date = date.AddDate(0, 0, 1)
	}
}

// HandlerRuleM обработчик для правила "m" (дни месяца)
func HandlerRuleM(now, date time.Time, parts []string) (string, error) {
	if len(parts) < 2 {
		return "", fmt.Errorf("missing days number for 'm' rule")
	}
	days, err := ParseMonthDays(parts[1])
	if err != nil {
		return "", err
	}
	months := []int{}
	if len(parts) >= 3 {
		months, err = ParseMonths(parts[2])
		if err != nil {
			return "", err
		}
	}

	date = date.AddDate(0, 0, 1)
	for {
		monthOk := len(months) == 0
		for _, month := range months {
			if int(date.Month()) == month {
				monthOk = true
				break
			}
		}
		if !monthOk {
			date = date.AddDate(0, 0, 1)
			continue
		}
		if ValidDayOfMonth(date, days) && date.After(now) {
			return date.Format(dateFormat), nil
		}
		date = date.AddDate(0, 0, 1)
	}
}

// NextDateHandler обрабатывает GET-запросы /api/nextdate
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

	nextD, err := NextDate(now, date, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte(nextD))
}

// PareWeekdays преобразует строку в мапу дней недели, для быстрой проверки
func ParseWeekdays(weekdays string) (map[time.Weekday]bool, error) {
	parts := strings.Split(weekdays, ",")
	result := make(map[time.Weekday]bool)
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 1 || num > 7 {
			return nil, fmt.Errorf("wrong weekday: %s", part)
		}
		var weekday time.Weekday
		if num == 7 {
			weekday = time.Sunday
		} else {
			weekday = time.Weekday(num)
		}
		result[weekday] = true
	}
	return result, nil
}

// ValidDayOfMonth проверяет подошёл ли день месяца под правило
func ValidDayOfMonth(date time.Time, days []int) bool {
	monthDay := date.Day()
	lastDay := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, date.Location()).Day()
	for _, day := range days {
		if day > 0 && monthDay == day {
			return true
		}
		if day < 0 && lastDay+1+day == monthDay {
			return true
		}
	}
	return false
}

// ParseMonthDays преобразует строку в слайс чисел
func ParseMonthDays(monthDays string) ([]int, error) {
	parts := strings.Split(monthDays, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("wrong day: %s", part)
		}
		if num > 0 && num <= 31 {
			result = append(result, num)
			continue
		}
		if num == -1 || num == -2 {
			result = append(result, num)
			continue
		}
		return nil, fmt.Errorf("wrong day: %s", part)
	}
	return result, nil
}

// ParseMonths преобразует строку в слайс чисел
func ParseMonths(allMonths string) ([]int, error) {
	parts := strings.Split(allMonths, ",")
	result := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num < 1 || num > 12 {
			return nil, fmt.Errorf("wrong month: %s", part)
		}
		result = append(result, num)
	}
	return result, nil
}
