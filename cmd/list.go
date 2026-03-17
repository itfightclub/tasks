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
	Args:  cobra.ExactArgs(0), // Prevent user error if typing more
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

		// Get the all flag
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return fmt.Errorf("failed to get all flag; %w", err)
		}

		// Display tasks
		taskList.ListTasks(all)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("all", "a", false, "List all tasks, including completed")
}
