package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/DylanMeeks/pokedexcli/internal/pokecache"
	"github.com/DylanMeeks/pokedexcli/internal/pokiapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*commandConfig) error
}

type commandConfig struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
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
		description: "Displays the names of next 20 location areas",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "Displays the names of the previous 20 location areas",
		callback:    commandMapb,
	},
}

func main() {

	sc := bufio.NewScanner(os.Stdin)

	config := commandConfig{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
		Cache:    pokecache.NewCache(time.Second * 2),
	}

	for {
		fmt.Print("Pokedex > ")
		sc.Scan()
		input := sc.Text()
		input = strings.ToLower(input)
		words := strings.Fields(input)

		if command, exists := registry[words[0]]; exists {
			if err := command.callback(&config); err != nil {
				fmt.Println(err)
			}

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
	cachedLocs, ok := config.Cache.Get(config.Next)
    var locs pokeapi.LocationReqRes
	// var locs []struct {
	// 	Name string `json:"name"`
	// 	URL  string `json:"url"`
	// }
	if !ok {
		res, err := pokeapi.GetLocations(config.Next)
		if err != nil {
			return err
		}
        locs = res
	} else {
		err := json.Unmarshal(cachedLocs, &locs)
		if err != nil {
			return err
		}
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	data, err := json.Marshal(locs)
    if err != nil {
        return err
    }
	config.Cache.Add(config.Next, data)

	config.Previous = config.Next
	config.Next = locs.Next
	return nil
}

func commandMapb(config *commandConfig) error {
	cachedLocs, ok := config.Cache.Get(config.Previous)
    var locs pokeapi.LocationReqRes
	// var locs []struct {
	// 	Name string `json:"name"`
	// 	URL  string `json:"url"`
	// }
	if !ok {
		res, err := pokeapi.GetLocations(config.Previous)
		if err != nil {
			return err
		}
        locs = res
	} else {
		err := json.Unmarshal(cachedLocs, &locs)
		if err != nil {
			return err
		}
	}

	for _, loc := range locs.Results {
		fmt.Println(loc.Name)
	}

	data, err := json.Marshal(locs)
    if err != nil {
        return err
    }
	config.Cache.Add(config.Previous, data)

	config.Next = config.Previous
	config.Previous = locs.Previous
	return nil
}
