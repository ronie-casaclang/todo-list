package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Clear the console screen
func clearScreen() {
	var clearCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default:
		clearCmd = exec.Command("clear")
	}
	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
}

// Pause
func pause() {
	fmt.Println("Press any key to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	clearScreen()
}
