/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

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

		file, err := os.Open("task.txt")
		if err != nil {
			fmt.Println("Error while loading tasks:", err)
			return
		}

		var tasks []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			tasks = append(tasks, scanner.Text())
		}
		file.Close()

		if taskID > len(tasks) {
			fmt.Println("Please enter a valid task id")
			return
		}

		completedTask := tasks[taskID-1]
		tasks = append(tasks[:taskID-1], tasks[taskID:]...)

		_ = os.WriteFile("task.txt", []byte(""), 0644)
		file, _ = os.OpenFile("task.txt", os.O_CREATE|os.O_WRONLY, 0644)
		for _, t := range tasks {
			_, _ = file.WriteString(t + "\n")
		}
		file.Close()

		file, _ = os.OpenFile("completed.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		defer file.Close()
		_, _ = file.WriteString(completedTask + "\n")
		fmt.Println("Task:", completedTask, "marked as completed.")
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
