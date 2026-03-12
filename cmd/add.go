/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
cmd/add.go
*/

package cmd

import (
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
		newTask := internal.Task{
			Name:    args[0],
			Created: time.Now(),
			Done:    false,
		}

		// TODO - use common access config file (csv) instead of making new
		taskList := internal.TaskList{}
		err := taskList.AddTask(newTask)
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
		taskList.ViewTasks()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
