/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
internal/tasks.go
*/

package internal

import (
	"fmt"
	"time"
)

const MaxTaskNameLength = 60 // Maximum allowed length for task names

type Task struct {
	Name    string    `csv:"Name"`
	Created time.Time `csv:"Created"`
	Done    bool      `csv:"Done"`
}

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) AddTask(newTask Task) error {
	if len(newTask.Name) > MaxTaskNameLength {
		return fmt.Errorf("task name exceeds maximum length of %d characters", MaxTaskNameLength)
	}
	tl.Tasks = append(tl.Tasks, newTask)
	return nil
}

func (tl *TaskList) MarkDone(index int) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks[index].Done = true
	}
}

func (tl *TaskList) DeleteTask(index int) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
	}
}

func (tl TaskList) ListTasks(all bool) {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	// Print header
	fmt.Printf("%-5s %-60s %-20s %-5s\n", "ID", "Task", "Created", "Done")

	// Iterate tasks
	for i, t := range tl.Tasks {
		if t.Done && !all {
			continue
		}
		createdDuration := durationToString(time.Since(t.Created))
		fmt.Printf("%-5d %-60s %-20s %-5t\n", i+1, t.Name, createdDuration, t.Done)
	}
}

// durationToString returns a nice, human readable, relative time
func durationToString(duration time.Duration) string {
	seconds := int(duration.Seconds())
	switch {
	case seconds < 60: // one minute
		return "a few seconds ago"
	case seconds < 120: // two minutes
		return "a minute ago"
	case seconds < 3600: // one hour
		return fmt.Sprintf("%d minutes ago", seconds/60)
	case seconds < 7200: // two hours
		return "an hour ago"
	case seconds < 86400: // 24 hours
		return fmt.Sprintf("%d hours ago", seconds/3600)
	default:
		return fmt.Sprintf("%d days ago", seconds/86400)
	}
}
