package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Main function
func main() {

	//Get user choice
	reader := bufio.NewReader(os.Stdin)
	for {
		clearScreen()
		fmt.Println("_______ ____ ______ ____")
		fmt.Println("|   |  || __||   | || | |")
		fmt.Println("| |  | || __|| | | || | |")
		fmt.Println("|_|__|_||___||_|___||___|")
		fmt.Println("< Todo List Application >")
		fmt.Println()
		fmt.Println("1. Add Task")
		fmt.Println("2. View List")
		fmt.Println("3. Edit Task")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")
		fmt.Print("\nSelect choice : ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			addTask()
		case "2":
			viewList()
		case "3":
			editTask()
		case "4":
			deleteTask()
		case "5":
			if exitApp() {
				pause()
				return
			}
		default:
			clearScreen()
			fmt.Println("Invalid Input!")
			pause()
		}
	}
}
