package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Dev-Jace/pokedex/internal/pokecache"
)

type Config struct {
	prev_URL  string
	next_URL  string
	cachePntr *pokecache.Cache
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	URL_config := Config{
		prev_URL: "none",
		next_URL: "none",
	}
	//make 'cache'
	cache := pokecache.NewCache(5 * time.Second)
	URL_config.cachePntr = cache

	for {
		fmt.Print("Pokedex > ") //, scanner.Text()
		scanner.Scan()
		usrInput := cleanInput(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}

		commandName := usrInput[0]
		var supCommand string
		if len(usrInput) > 1 {
			supCommand = usrInput[1]
		}
		commands := getCommands()
		_, exists := commands[commandName]
		if exists {
			err := commands[commandName].callback(&URL_config, supCommand)
			if err != nil {
				fmt.Printf("encountered error: '%v'\n", err)
			}
		} else {
			fmt.Printf("command '%v' not found\n", commandName)
		}

	} //end for loop
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

func getCommands() map[string]cliCommand {

	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a map location's Pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "try and catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "get the stats of a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "get list of caught pokemon",
			callback:    commandPokedex,
		},
	}
}
