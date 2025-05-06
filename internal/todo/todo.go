package todo

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func AddNewTask(task string) {
	file, err := os.OpenFile("task.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while creating a new task.", err)
		return
	}
	defer file.Close()

	task = strings.Trim(task, " ")
	_, err = file.WriteString(task + "\n")
	if err != nil {
		fmt.Println("Error while creating a new task:", err)
		return
	}
	fmt.Println("Created new task:", task)
}

func CompleteTask(taskID int) {
	tasks := loadTasks(false)
	if len(tasks) == 0 {
		fmt.Println("Please create a task first.")
		return
	}
	if taskID > len(tasks) {
		fmt.Println("Please enter a valid task-id.")
		return
	}

	err := os.WriteFile("task.txt", []byte(""), 0644)
	if err != nil {
		panic("Error while removing tasks.")
	}

	completedTask := tasks[taskID-1]
	tasks = slices.Delete(tasks, taskID-1, taskID-1)

	file, err := os.OpenFile("task.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while re-listing tasks.")
	}
	defer file.Close()

	for _, t := range tasks {
		_, err = file.WriteString(t + "\n")
		if err != nil {
			panic("Error while re-listing tasks.")
		}
	}

	if file, err := os.OpenFile("completed.txt", os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		panic("Error while marking task as completed.")
	} else {
		file.WriteString(completedTask)
		file.Close()
	}

	fmt.Println("Successfully marked task as completed:", completedTask)
}

func ListTasks(showCompleted bool) {
	tasks := loadTasks(showCompleted)
	if len(tasks) == 0 {
		if showCompleted {
			fmt.Println("Please mark a task as complete first. Usage: clitodo add task-id")
		} else {
			fmt.Println("Please create a task first. Usage: clitodo add `task-name`")
		}
		return
	}
	if showCompleted {
		fmt.Println("Completed TODO list:")
	} else {
		fmt.Println("Current TODO list: ")
	}
	for i, t := range tasks {
		fmt.Printf("%d. %s\n", i+1, t)
	}
}

func RemoveTask(taskID int) {
	tasks := loadTasks(false)
	if len(tasks) == 0 {
		fmt.Println("Please create a task first. Usage: clitodo add `task-name`")
		return
	}

	if taskID > len(tasks) {
		fmt.Println("Please enter a valid task-id.")
		return
	}

	err := os.WriteFile("task.txt", []byte(""), 0644)
	if err != nil {
		panic("Error while removing tasks.")
	}

	removedTask := tasks[taskID-1]
	tasks = slices.Delete(tasks, taskID-1, taskID-1)

	file, err := os.OpenFile("task.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error while re-listing tasks.")
	}
	defer file.Close()

	for _, t := range tasks {
		_, err = file.WriteString(t + "\n")
		if err != nil {
			panic("Error while re-listing tasks.")
		}
	}
	fmt.Println("Successfully removed task:", removedTask)
}

func loadTasks(showCompleted bool) []string {
	var file *os.File
	var err error

	if showCompleted {
		file, err = os.Open("completed.txt")
		if err != nil {
			panic("Error while loading completed tasks.")
		}
	} else {
		file, err = os.Open("task.txt")
		if err != nil {
			panic("Error while loading tasks.")
		}
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}
	return tasks
}
