/*
SPDX-License-Identifier: MIT
SPDX-FileCopyrightText: C 2026 https://github.com/itfightclub
cmd/version.go
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Meant to be the source of truth
const Version = "0.3.0-dev"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("tasks %s\n", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
