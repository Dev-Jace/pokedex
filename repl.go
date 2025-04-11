package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ") //, scanner.Text()
		scanner.Scan()
		usrInput := cleanInput(scanner.Text())
		if len(usrInput) == 0 {
			continue
		}
		commandName := usrInput[0]
		commands := getCommands()
		_, exists := commands[commandName]
		if exists {
			commands[commandName].callback()
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
	callback    func() error
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
	}
}
