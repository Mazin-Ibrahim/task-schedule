# Task Scheduler CLI

## Overview
Task Scheduler is a command-line application built with Golang that allows users to manage scheduled tasks. The project uses PostgreSQL as the database backend to store task details. It notifies users when the scheduled time for a task arrives.

## Features
- Add tasks with a name, description, and schedule.
- Update existing tasks.
- List all scheduled tasks.
- Delete tasks by ID.
- Schedule tasks and receive notifications when they are due.


### Add a Task
```sh
go run *.go add --name="task name" --description="task  description" --schedule="0 9 * * *"
```

### Update a Task
```sh
go run *.go update --id="1" --name="task 2 updated" --description="task updated" --schedule="0 9 * * *"
```

### List All Tasks
```sh
go run *.go list
```

### Delete a Task
```sh
go run *.go delete --id=3
```

### Schedule a Task
```sh
go run *.go schedule --id=2
```

## Notifications
The application will notify you when the scheduled time for a task arrives. Ensure notifications are properly configured in your system to receive alerts.



