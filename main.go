package main

import (
	"bufio"
	"fmt"
	"github.com/DylanMeeks/pokedexcli/internal/pokiapi"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*commandConfig) error
}

type commandConfig struct {
	Next     string
	Previous string
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
	"map": {
		name:        "map",
		description: "Displays the names of 20 location areas",
		callback:    commandMap,
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
			command.callback(new(commandConfig))
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	words := strings.Fields(text)
	return words
}

func commandExit(config *commandConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *commandConfig) error {
	fmt.Println(
		`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex`)
	return nil
}

func commandMap(config *commandConfig) error {
    GetLocationArea
}
