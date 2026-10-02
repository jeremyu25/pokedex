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
	callback    func(*config, ...string) error
}

func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	splitString := strings.Fields(text)
	var finalSlice []string

	for _, word := range splitString {
		finalSlice = append(finalSlice, strings.ToLower(word))
	}
	return finalSlice
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Displays 20 location areas in Pokemon World",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays 20 previous location areas in Pokemon World",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <name of location area>",
			description: "Displays list of Pokemon in location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <name of pokemon>",
			description: "Attempts to catch a specified Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <name of pokemon>",
			description: "Inspect a caught Pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List the name of all Pokemon in user's pokedex",
			callback:    commandPokedex,
		},
	}
}

func startRepl(cfg *config) {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		scanner.Scan()
		if scanner.Err() != nil {
			os.Exit(0)
		}
		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}
		command, found := cfg.commands[cleanedInput[0]]
		if found {
			if err := command.callback(cfg, cleanedInput[1:]...); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
