/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "The add subcommand will add a new task for todo list.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if task == "" {
			fmt.Println("Please consider providing a task name")
			return
		}

		file, err := os.OpenFile("task.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error while saving a task:", err)
			return
		}
		defer file.Close()
		task = strings.Trim(task, " ")
		_, err = file.WriteString(task + "\n")
		if err != nil {
			fmt.Println("Error while saving a task:", err)
			return
		} else {
			fmt.Println("Task added:", task)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
