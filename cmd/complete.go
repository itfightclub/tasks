/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
cmd/complete.go
*/

package cmd

import (
	"fmt"
	"strconv"

	"github.com/itfightclub/tasks/internal"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task complete",
	Long:  "Mark a task as complete by specifying its ID number",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the tasks file path from the persistent flag
		tasksFile, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("failed to get config flag: %w", err)
		}

		// Parse the task ID
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %s. ID must be a number", args[0])
		}

		// Load existing tasks from CSV
		taskList, err := internal.LoadTasks(tasksFile)
		if err != nil {
			return fmt.Errorf("failed to load tasks: %w", err)
		}

		// Convert to 0-based index (CLI uses 1-based indexing)
		index := id - 1

		// Check if tasks exists
		if index < 0 || index >= len(taskList.Tasks) {
			return fmt.Errorf("task ID %d not found. Valid IDs are 1-%d", id, len(taskList.Tasks))
		}

		// Check if already completed
		if taskList.Tasks[index].Done {
			return fmt.Errorf("task %d is already marked as complete", id)
		}

		// Mark task as done
		taskList.MarkDone(index)

		// Save updated task list back to CSV
		if err := internal.SaveTasks(tasksFile, taskList); err != nil {
			return fmt.Errorf("failed to save tasks: %w", err)
		}

		// Show confirmation
		fmt.Printf("Task %d marked as complete: %s\n", id, taskList.Tasks[index].Name)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
