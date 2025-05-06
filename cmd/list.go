/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ryszhio/clitodo/internal/todo"
	"github.com/spf13/cobra"
)

var showCompleted bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "The add sub command will list the current task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		todo.ListTasks(showCompleted)
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showCompleted, "completed", "c", false, "Shows only completed tasks")
	rootCmd.AddCommand(listCmd)
}
