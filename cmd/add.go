/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/ryszhio/clitodo/internal/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "The add subcommand will add a new task for todo list.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		todo.AddNewTask(task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
