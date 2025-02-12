package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {

	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		input := sc.Text()
		input = strings.ToLower(input)
		words := strings.Fields(input)
        
        if command, exists := commands[words[0]]; exists {
            command.callback()
        } else {
            fmt.Println("Unknown command")
        }
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
