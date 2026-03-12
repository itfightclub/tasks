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
	// ISSUE: this doesn't actually print with verbose
	// PreRun: func(cmd *cobra.Command, args []string) {
	// 	// Create the tasks file before running the command
	// 	if err := createTasksFile(tasksFile); err != nil {
	// 		println("Error creating tasks file:", err.Error())
	// 		os.Exit(1) // Exit if there is an error in creating the tasks file
	// 	}
	// },
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

	// ISSUE: this doesn't actually print with verbose
	// Create the tasks file in the init()
	// if err := createTasksFile(tasksFile); err != nil {
	// 	println("Error creating tasks file:", err.Error())
	// 	os.Exit(1) // Exit if there is an error in creating the tasks file
	// }
}

// TODO - find out where to call this so verbose shows it being created. see 'ISSUE:' comments
// creates the tasksFile if it doesn't exist
func createTasksFile(filename string) error {
	// Check if the file exists
	if _, err := os.Stat(filename); err == nil {
		if verbose {
			println("Task file found:", filename)
		}
		return nil // File exists, no need to create
	} else if os.IsNotExist(err) {
		// File does not exist, create directories and file
		if verbose {
			println("Tasks file not found. Creating:", filename)
		}

		// Create the directory if it doesn't exist
		dir := filepath.Dir(filename)
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err // Return an error if directory creation fails
		}

		// Create the tasks file
		file, err := os.Create(filename)
		if err != nil {
			return err // Return an error if file creation fails
		}
		defer file.Close() // Ensure the file gets closed

		if verbose {
			println("Tasks file created successfully:", filename)
		}
		return nil // Successfully created dir and file

	} else {
		return err // Return other errors
	}
}
