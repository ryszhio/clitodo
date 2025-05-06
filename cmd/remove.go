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

		file, err := os.Open("task.txt")
		if err != nil {
			fmt.Println("Error while fetching tasks.", err)
			return
		}

		var tasks []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			tasks = append(tasks, scanner.Text())
		}
		file.Close()

		if taskID > len(tasks) {
			fmt.Println("Please provide a valid task id.")
			return
		}

		removedTask := tasks[taskID-1]
		tasks = append(tasks[:taskID-1], tasks[taskID:]...)

		err = os.WriteFile("task.txt", []byte(""), 0644)
		if err != nil {
			fmt.Println("Error while removing task:", err)
			return
		}
		file, _ = os.OpenFile("task.txt", os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()
		for _, t := range tasks {
			_, _ = file.WriteString(t + "\n")
		}
		fmt.Println("Removed Task: ", removedTask)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
