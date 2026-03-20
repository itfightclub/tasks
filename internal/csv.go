/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
internal/csv.go
*/

package internal

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// LoadTasks reads tasks from a CSV file and returns a TaskList.
// If the file doesn't exist, it will create it.
func LoadTasks(filename string) (*TaskList, error) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist yet, return empty task list
			return &TaskList{Tasks: []Task{}}, nil
		}
		return nil, fmt.Errorf("failed to open tasks file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Read header row
	header, err := reader.Read()
	if err != nil {
		if err == io.EOF {
			// Empty file, return empty task list
			return &TaskList{Tasks: []Task{}}, nil
		}
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	// Validate header
	expectedHeader := []string{"Name", "Created", "Done"}
	if len(header) != len(expectedHeader) {
		return nil, fmt.Errorf("invalid CSV header: expected %d columns, got %d", len(expectedHeader), len(header))
	}
	for i, col := range expectedHeader {
		if strings.TrimSpace(header[i]) != col {
			return nil, fmt.Errorf("invalid CSV header: expected '%s' at position %d, got '%s'", col, i, header[i])
		}
	}

	var tasks []Task
	lineNum := 1 // Start at 2 since header is line 1

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read CSV line %d: %w", lineNum+1, err)
		}

		if len(record) != 3 {
			return nil, fmt.Errorf("invalid record at line %d: expected 3 fields, got %d", lineNum+1, len(record))
		}

		// Parse Created timestamp
		created, err := time.Parse(time.RFC3339, strings.TrimSpace(record[1]))
		if err != nil {
			return nil, fmt.Errorf("invalid timestamp at line %d: %w", lineNum+1, err)
		}

		// Parse Done boolean
		doneStr := strings.ToLower(strings.TrimSpace(record[2]))
		var done bool
		switch doneStr {
		case "true", "1", "yes":
			done = true
		case "false", "0", "no", "":
			done = false
		default:
			return nil, fmt.Errorf("invalid boolean value at line %d: '%s'", lineNum+1, doneStr)
		}

		task := Task{
			Name:    strings.TrimSpace(record[0]),
			Created: created,
			Done:    done,
		}

		tasks = append(tasks, task)
		lineNum++
	}

	return &TaskList{Tasks: tasks}, nil
}

// SaveTasks writes a TaskList to a CSV file
func SaveTasks(filename string, taskList *TaskList) error {
	// Ensure directory exists
	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create tasks file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"Name", "Created", "Done"}); err != nil {
		return fmt.Errorf("")
	}

	// Write tasks
	for i, task := range taskList.Tasks {
		record := []string{
			task.Name,
			task.Created.Format(time.RFC3339),
			strconv.FormatBool(task.Done),
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write task %d: %w", i+1, err)
		}
	}

	return nil
}
