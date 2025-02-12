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

var registry map[string]cliCommand = map[string]cliCommand{
    "exit": {
        name:        "exit",
        description: "Exit the Pokedex",
        callback:    commandExit,
    },
    "help": {
        name:        "help",
        description: "Displays a help message",
        callback:    commandHelp,
    },
}

func main() {

	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		input := sc.Text()
		input = strings.ToLower(input)
		words := strings.Fields(input)
        
        if command, exists := registry[words[0]]; exists {
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

func commandHelp() error {
    fmt.Println(
`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
    return nil
}
