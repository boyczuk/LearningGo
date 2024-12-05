package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []string

func todolist() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Select an option:")
		fmt.Println("(add) Add task\n(complete) Complete task\n(show) Display tasks\n(exit) Exit")
		userInput, _ := reader.ReadString('\n')
		input := strings.TrimSpace(strings.ToLower(userInput))

		if input == "exit" {
			break
		} else if input == "add" {
			addTask()
		} else if input == "complete" {
			completeTask()
		} else if input == "show" {
			showTasks()
		} else {
			fmt.Println("Not valid choice")
		}
	}

}

func addTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add task: ")
	input, _ := reader.ReadString('\n')
	tasks = append(tasks, input)
}

func completeTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Complete task: ")
	input, _ := reader.ReadString('\n')
	index := -1
	for i, task := range tasks {
		if task == input {
			index = i
			break
		}
	}

	if index != -1 {
		var removedTask = tasks[index]
		tasks = append(tasks[:index], tasks[index+1:]...)

		fmt.Println("Removed task: " + removedTask)
	}

}

func showTasks() {
	fmt.Println("Show tasks")
	for index, task := range tasks {
		index = index + 1
		fmt.Println(strconv.Itoa(index) + ". " + task)
	}
}

func main() {
	todolist()
}
