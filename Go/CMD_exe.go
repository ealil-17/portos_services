package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Prompt the user for a command
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the command: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	// Parse the command and arguments
	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")
	if len(parts) == 0 {
		fmt.Println("No command provided")
		return
	}

	command := parts[0]
	args := parts[1:]

	// Create a command instance
	cmd := exec.Command(command, args...)

	// Set up the command to print logs in real time
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	fmt.Println("Command executed successfully")
}
