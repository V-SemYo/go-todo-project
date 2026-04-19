package db

import (
	"database/sql"
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

// GetTask возвращает задачу по её ID (или err)
func GetTask(id string) (*Task, error) {
	task := &Task{}
	query := `SELECT id, date, title, comment, repeat FROM scheduler WHERE id = ?`
	err := DB.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("task not found")
		}
		return nil, fmt.Errorf("getting task error: %w", err)
	}
	return task, nil
}

// UpdateTask обновляет существующую задачу, если задачи нет вернёт err
func UpdateTask(task *Task) error {
	query := `UPDATE scheduler SET date = ?, title = ?, comment = ?, repeat = ? WHERE id = ?`
	res, err := DB.Exec(query, task.Date, task.Title, task.Comment, task.Repeat, task.ID)
	if err != nil {
		return fmt.Errorf("updating task error: %w", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting rows affected error: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("task not found")
	}
	return nil
}

// DeleteTask удаляет задачу по указанному id
func DeleteTask(id string) error {
	query := `DELETE FROM scheduler WHERE id = ?`
	res, err := DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("deleting task error: %w", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting rows affected error: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("task not found")
	}
	return nil
}

// UpdateTaskDate обновляет дату указанной задачи
func UpdateTaskDate(id, newDate string) error {
	query := `UPDATE scheduler SET date = ? WHERE id = ?`
	res, err := DB.Exec(query, newDate, id)
	if err != nil {
		return fmt.Errorf("updating date error: %w", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("getting rows affected error: %w", err)
	}
	if count == 0 {
		return fmt.Errorf("task not found")
	}
	return nil
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
