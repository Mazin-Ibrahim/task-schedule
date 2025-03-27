package main

import (
	"database/sql"
	"errors"
	"time"
)

type Task struct {
	ID          int64
	Name        string
	Description string
	Schedule    string
	IsSchedule  bool
	IsRecurring bool
	CreatedAt   time.Time
}

func (t *Task) create(db *sql.DB, name string, descriotion string, schedule string) (*Task, error) {
	query := `
		INSERT INTO tasks (name,description,schedule)
		VALUES ($1,$2,$3)
		RETURNING id,name
	`
	var task Task

	err := db.QueryRow(query, name, descriotion, schedule).Scan(&task.Name, &task.Description)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (t *Task) update(db *sql.DB, id int64, name string, description string, schedule string) error {

	task, err := t.getTask(db, id)
	if err != nil {
		return err
	}
	if name == "" {
		name = task.Name
	}
	if description == "" {
		description = task.Description
	}

	if schedule == "" {
		schedule = task.Schedule
	}

	query := `
	    UPDATE tasks SET name=$1,description=$2,schedule=$3 WHERE id=$4
	`
	res, err := db.Exec(query, name, description, schedule, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("the task not found")
	}
	return nil
}

func (t *Task) getTask(db *sql.DB, id int64) (*Task, error) {

	query := `SELECT id,name,created_at,schedule,is_schedule FROM tasks WHERE id=$1`
	var task Task

	err := db.QueryRow(query, id).Scan(&task.ID, &task.Name, &task.CreatedAt, &task.Schedule, &task.IsSchedule)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *Task) getTasks(db *sql.DB) ([]Task, error) {

	query := `SELECT id,name,schedule,created_at FROM tasks`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Name, &task.Schedule, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *Task) delete(db *sql.DB, id int64) error {

	query := `DELETE FROM tasks WHERE id= $1`

	res, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("the task not found")
	}
	return nil
}

func (t *Task) setTaskSchedule(db *sql.DB, id int64) error {

	query := `
	    UPDATE tasks SET is_schedule = $1 WHERE id = $2 
	`

	res, err := db.Exec(query, true, id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("the task not found")
	}
	return nil
}
