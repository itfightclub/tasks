/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
internal/tasks.go
*/

package internal

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

const MaxTaskNameLength = 80 // Maximum allowed length for task names

// Task is the definition of a task
type Task struct {
	Name    string    `csv:"Name"`
	Created time.Time `csv:"Created"`
	Done    bool      `csv:"Done"`
}

// TaskList is a list of tasks
type TaskList struct {
	Tasks []Task
}

// AddTask adds a task to a TaskList
func (tl *TaskList) AddTask(newTask Task) error {
	if len(newTask.Name) > MaxTaskNameLength {
		return fmt.Errorf("task name exceeds maximum length of %d characters", MaxTaskNameLength)
	}
	tl.Tasks = append(tl.Tasks, newTask)
	return nil
}

// MarkDone marks a task as done
func (tl *TaskList) MarkDone(index int) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks[index].Done = true
	}
}

// DeleteTask deletes a task
func (tl *TaskList) DeleteTask(index int) {
	if index >= 0 && index < len(tl.Tasks) {
		tl.Tasks = append(tl.Tasks[:index], tl.Tasks[index+1:]...)
	}
}

// ListTasks lists all the tasks in a TaskList
func (tl TaskList) ListTasks(all bool) {
	if len(tl.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	// Create tab writer
	w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)

	// Print header
	fmt.Fprintln(w, "ID\tTask\tCreated\tDone")

	// Iterate tasks
	for i, t := range tl.Tasks {
		if t.Done && !all {
			continue
		}
		createdDuration := durationToString(time.Since(t.Created))
		fmt.Fprintf(w, "%d\t%s\t%s\t%t\n", i+1, t.Name, createdDuration, t.Done)
	}

	// Flush the writer to ensure the output is displayed
	if err := w.Flush(); err != nil {
		fmt.Println("Error flushing writer:", err)
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
