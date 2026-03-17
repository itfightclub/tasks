/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
cmd/root.go
*/

package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	tasksFile string
	verbose   bool
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "A Todo CLI App",
	Long:  `A simple CLI to-do app`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	defaultTasksFile := filepath.Join(os.Getenv("HOME"), ".tasks", "tasks.csv")
	rootCmd.PersistentFlags().StringVar(&tasksFile, "config", defaultTasksFile, "config file")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "enable verbosity")
}
