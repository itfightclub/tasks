/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
cmd/delete.go
*/

package cmd

import (
	"fmt"
	"strconv"

	"github.com/itfightclub/tasks/internal"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a task permanently",
	Long:  "Permanently remove a task from the list by specifying its ID number",
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

		// Load existing task from CSV
		taskList, err := internal.LoadTasks(tasksFile)
		if err != nil {
			return fmt.Errorf("failed to load tasks: %w", err)
		}

		// Convert to 0-based index (CLI uses 1-based indexing)
		index := id - 1

		// Check if tasks exists
		if index < 0 || index >= len(taskList.Tasks) {
			return fmt.Errorf("task ID %d no found. Valid IDs are 1-%d", id, len(taskList.Tasks))
		}

		// Get task name for confirmation message
		taskName := taskList.Tasks[index].Name

		// Delete the task
		taskList.DeleteTask(index)

		// Save updated task list back to CSV
		if err := internal.SaveTasks(tasksFile, taskList); err != nil {
			return fmt.Errorf("failed to save tasks: %w", err)
		}

		// Show confirmation
		fmt.Printf("Task %d deleted: %s\n", id, taskName)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
