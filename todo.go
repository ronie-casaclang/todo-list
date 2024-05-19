package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Structure for todo item
type Todo struct {
	ID   int
	Task string
}

// Global variable to store the list
var todoList []Todo
var nextID int = 1

// Add a new task
func addTask() {
	clearScreen()
	fmt.Println("____ ___ ___")
	fmt.Println("| _ ||  \\|  \\")
	fmt.Println("|   || | | | |")
	fmt.Println("|_|_||__/|__/")

	//Get user input
	fmt.Print("Enter a task : ")
	reader := bufio.NewReader(os.Stdin)
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	//Add into slice then increment the ID for next task
	todoList = append(todoList, Todo{ID: nextID, Task: task})
	nextID++
	fmt.Println("Task Added!")
	pause()
}

// Display list of task
func viewList() {
	clearScreen()
	fmt.Println("__  __ ____ ____")
	fmt.Println("| | | || __||_ _|")
	fmt.Println("| |_| ||__ | | |")
	fmt.Println("|___|_||___| |_|")

	if len(todoList) == 0 {
		fmt.Println("Empty List!")
	} else {
		fmt.Println("Todo List :")
		for i := 0; i < len(todoList); i++ {
			fmt.Printf("%v. %v \n", todoList[i].ID, todoList[i].Task)
		}
		fmt.Println("------------------------------------------------")
	}
	pause()
	return
}

// Edit a task
func editTask() {
	clearScreen()
	fmt.Println("____ ____ __ _______")
	fmt.Println("|  _||   \\| ||_   _|")
	fmt.Println("|  _|| |  | |  | |")
	fmt.Println("|___||___/|_|  |_|")

	if len(todoList) == 0 {
		fmt.Println("Empty List!")
	} else {
		fmt.Print("Enter task ID : ")
		var inputID int
		fmt.Scanf("%d", &inputID)
		fmt.Scanln()
		for i, todo := range todoList {
			if todo.ID == inputID {
				fmt.Println("Current value   : ", todoList[i].Task)
				fmt.Print("Enter new value : ")

				//get user input
				reader := bufio.NewReader(os.Stdin)
				newTask, _ := reader.ReadString('\n')
				newTask = strings.TrimSpace(newTask)
				todoList[i].Task = newTask
				fmt.Println("Task updated!")
				pause()
				return
			}
		}
		fmt.Println("Task not found!")
	}
	pause()
	return
}

// Delete a task
func deleteTask() {
	clearScreen()
	fmt.Println("____ ____ ______ ____ __   __ ____")
	fmt.Println("| _ \\|  _||  |  ||   |\\ \\ / / | _|")
	fmt.Println("|   ||  _|| | | || | | \\   /  | _|")
	fmt.Println("|_|_||___||_|_|_||___|  \\_/   |__|")

	if len(todoList) == 0 {
		fmt.Println("Empty List!")
	} else {
		fmt.Print("Enter task ID : ")
		var inputID int
		fmt.Scanf("%d", &inputID)
		fmt.Scanln()
		for i, todo := range todoList {
			if todo.ID == inputID {
				todoList = append(todoList[:i], todoList[i+1:]...)
				fmt.Println("Task deleted!")
				pause()
				return
			}
		}
		fmt.Println("Task not found!")
	}
	pause()
	return
}

// Exit Application
func exitApp() bool {
	clearScreen()
	fmt.Println("____ __  ____________")
	fmt.Println("|  _|\\ \\/ /| ||_   _|")
	fmt.Println("|  _| /  \\ | |  | |")
	fmt.Println("|___|/_/\\_\\|_|  |_|")
	fmt.Print("Are you sure you want to exit? [Y/N] : ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "Y" || choice == "y" {
		return true
	}
	return false
}
