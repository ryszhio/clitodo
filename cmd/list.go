/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var showCompleted bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "The add sub command will list the current task",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var file *os.File
		var err error
		if showCompleted {
			file, err = os.Open("completed.txt")
			if err != nil {
				fmt.Println("There is no completed tasks.")
				return
			}
		} else {
			file, err = os.Open("task.txt")
			if err != nil {
				fmt.Println("Error while loading task:", err)
				return
			}
		}
		defer file.Close()

		var tasks []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			tasks = append(tasks, scanner.Text())
		}
		if showCompleted {
			fmt.Println("Completed Tasks:")
		} else {
			fmt.Println("Current Tasks:")
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task)
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&showCompleted, "completed", "c", false, "Shows only completed tasks")
	rootCmd.AddCommand(listCmd)
}
