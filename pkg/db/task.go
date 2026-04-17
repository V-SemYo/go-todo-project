package db

import (
	"fmt"
	"time"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

// AddTask добавляет новую задачу в таблицу scheduler и возвращает ID(или err)
func AddTask(task *Task) (int64, error) {
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES (?, ?, ?, ?)`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat)
	if err != nil {
		return 0, fmt.Errorf("inserting error: %w", err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("getting last insert error: %w", err)
	}
	return id, nil
}

// Tasks возвращает список задач, отсортированных по дате из таблицы scheduler
func Tasks(limit int, search string) ([]*Task, error) {
	var query string
	var params []any

	if search == "" {
		query = `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date LIMIT ?`
		params = []any{limit}
	} else {
		if tDate, err := time.Parse("02.01.2006", search); err == nil {
			dateStr := tDate.Format("20060102")
			query = `SELECT id, date, title, comment, repeat FROM scheduler WHERE date = ? ORDER BY date LIMIT ?`
			params = []any{dateStr, limit}
		} else {
			searchPattern := "%" + search + "%"
			query = `SELECT id, date, title, comment, repeat FROM scheduler WHERE title LIKE ? OR comment LIKE ? ORDER BY date LIMIT ?`
			params = []any{searchPattern, searchPattern, limit}
		}
	}
	rows, err := DB.Query(query, params...)
	if err != nil {
		return nil, fmt.Errorf("tasks query error: %w", err)
	}
	defer rows.Close()

	var tasks = []*Task{}
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, fmt.Errorf("tasks scan error: %w", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
