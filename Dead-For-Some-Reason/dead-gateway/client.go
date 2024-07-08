package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("$ ")
        // Read the keyboad input.
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
        }

        // Handle the execution of the input.
        if err = execInput(input); err != nil {
            fmt.Fprintln(os.Stderr, err)
        }
    }
    os.Exit(0)
}


func handleInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")
	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")
	
	// Check for built-in commands.
	switch args[0] {
	case "connect":
		establishConnection()
	case "help":
		printHelp()	
	case "exit":
    		os.Exit(0)
	}
}

func printHelp() {
	fmt.Println("USAGE:\t<command> <arguments>")
	fmt.Println("\t\tconnect <teamserver address>")
	fmt.Println("\t\thelp: get some help")
	fmt.Println("\t\texit to exit")
}

func establishConnection() error {

}
