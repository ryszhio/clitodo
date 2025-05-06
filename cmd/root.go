/*
Copyright © 2025 Rishab Karki rishabkarki78@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "clitodo",
	Version: "0.0.1",
	Short:   "Manage your todo list",
	Long: `
clitodo is a CLI based app as name suggest which can be
used to manage your task and stuffs from CLI itself.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
