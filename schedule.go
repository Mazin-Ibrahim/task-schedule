package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/robfig/cron"
)

func scheduleTask(db *sql.DB, id int64) error {
	task, err := (&Task{}).getTask(db, id)

	if err != nil {
		return err
	}
	if task.IsSchedule {
		return errors.New("the task is scheduled before")
	}
	task.setTaskSchedule(db, id)
	c := cron.New()
	c.AddFunc(task.Schedule, func() {
		fmt.Printf("Running task: %v\n %s\n", time.Now(), task.Name)
		// Add task logic here
	})
	c.Start()
	return nil
}
