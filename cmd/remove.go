/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/ryszhio/clitodo/internal/todo"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "The remove command will remove a certain task from task list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: remove task-id")
			return
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil || taskID < 1 {
			fmt.Println("Please provide a valid task id.")
			return
		}
		todo.RemoveTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
