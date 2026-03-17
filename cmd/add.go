/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
cmd/add.go
*/

package cmd

import (
	"fmt"
	"time"

	"github.com/itfightclub/tasks/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <task>",
	Short: "Add a new task (max length: 60 characters)",
	Long:  "Add a new task with the specified name. The task name must not exceed 60 characters",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the tasks file path from the persistent flag
		tasksFile, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("failed to get config file: %w", err)
		}

		// Get verbosity
		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return fmt.Errorf("failed to get verbosity flag: %v", verbose)
		}

		// Load existing tasks from CSV
		taskList, err := internal.LoadTasks(tasksFile)
		if err != nil {
			return fmt.Errorf("failed to load tasks: %w", err)
		}

		// Create new task
		newTask := internal.Task{
			Name:    args[0],
			Created: time.Now(),
			Done:    false,
		}

		// Add the new task
		if err := taskList.AddTask(newTask); err != nil {
			return fmt.Errorf("Failed to add task: %w", err)
		}

		// Save updated task list back to CSV
		if err := internal.SaveTasks(tasksFile, taskList); err != nil {
			return fmt.Errorf("failed to save tasks: %w", err)
		}

		// Show confirmation
		if verbose {
			fmt.Printf("Tasks added successfully: %s\n", newTask.Name)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
