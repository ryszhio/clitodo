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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "The complete command will mark a task as complete",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: complete task-id")
			return
		}
		taskID, err := strconv.Atoi(args[0])
		if err != nil || taskID < 1 {
			fmt.Println("Please enter a valid task id")
			return
		}
		todo.CompleteTask(taskID)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
