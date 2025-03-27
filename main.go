package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	db, err := connectDB()
	if err != nil {
		log.Fatalf("Faild to connect db,%v", err)
	}
	defer closeDB(db)

	taskCmd := flag.NewFlagSet("Task", flag.ExitOnError)
	name := taskCmd.String("name", "", "Name of the task")
	description := taskCmd.String("description", "", "Description of the task")
	schedule := taskCmd.String("schedule", "", "Schedule of the task (cron format)")
	id := taskCmd.String("id", "", "Id of task you want to make opration on it")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run *.go <command> [flags]")
		fmt.Println("Commands:")
		fmt.Println("  add - Add a new task")
		return
	}

	task := &Task{}
	switch os.Args[1] {
	case "add":
		taskCmd.Parse(os.Args[2:])
		if *name == "" || *schedule == "" {
			fmt.Println("Error: --name and --schedule are required")
			taskCmd.PrintDefaults()
			return
		}

		_, err := task.create(db, *name, *description, *schedule)
		if err != nil {
			log.Fatalf("Faild to add new task %v", err)
			break
		}
		fmt.Println("Task added successfully")
		return
	case "list":
		tasks, err := task.getTasks(db)
		if err != nil {
			log.Fatalf("Faild to fetch tasks %v", err)
			break
		}
		for _, task := range tasks {
			fmt.Printf("id : %d | name : %s", task.ID, task.Name)
			fmt.Println()
		}
		return

	case "delete":
		taskCmd.Parse(os.Args[2:])
		if *id == "" {
			fmt.Println("Error: --id and is required")
			taskCmd.PrintDefaults()
			return
		}
		parseId, err := strconv.ParseInt(*id, 10, 64)
		if err != nil {
			log.Fatalf("Faild to parse id %v", err)
			return
		}
		err = task.delete(db, parseId)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Task deleted successfuly")
		return
	case "update":
		taskCmd.Parse(os.Args[2:])
		if *id == "" {
			fmt.Println("Error: --id and is required")
			taskCmd.PrintDefaults()
			return
		}
		parseId, err := strconv.ParseInt(*id, 10, 64)
		if err != nil {
			log.Fatalf("Faild to parse id %v", err)
			return
		}
		err = task.update(db, parseId, *name, *description, *schedule)
		if err != nil {
			log.Fatalf("Faild to update task %v", err)
			return
		}
		fmt.Println("Task update successfuly")
		return

	case "one-task":
		taskCmd.Parse(os.Args[2:])
		if *id == "" {
			fmt.Println("Error: --id and is required")
			taskCmd.PrintDefaults()
			return
		}
		parseId, err := strconv.ParseInt(*id, 10, 64)
		if err != nil {
			log.Fatalf("Faild to parse id %v", err)
			return
		}
		task, err := task.getTask(db, parseId)
		if err != nil {
			log.Fatalf("Faild to fetch task %v", err)
			return
		}
		fmt.Printf("id : %d | name : %s", task.ID, task.Name)
		fmt.Println()
		return

	case "schedule":
		taskCmd.Parse(os.Args[2:])
		if *id == "" {
			fmt.Println("Error: --id and is required")
			taskCmd.PrintDefaults()
			return
		}
		parseId, err := strconv.ParseInt(*id, 10, 64)
		if err != nil {
			log.Fatalf("Faild to parse id %v", err)
			return
		}

		if err = scheduleTask(db, parseId); err != nil {
			log.Fatal(err)
			return
		}
		select {}

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}

}
