/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
cmd/list.go
*/

package cmd

import (
	"fmt"

	"github.com/itfightclub/tasks/internal"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List out all tasks status",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get the tasks file path from the persistent flag
		tasksFile, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("failed to get config flag: %w", err)
		}

		// Load tasks from CSV
		taskList, err := internal.LoadTasks(tasksFile)
		if err != nil {
			return fmt.Errorf("failed to load tasks: %w", err)
		}

		// Display tasks
		taskList.ListTasks()

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
