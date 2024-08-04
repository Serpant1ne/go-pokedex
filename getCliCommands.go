package main

import "github.com/Serpant1ne/go-pokedex/internal/pokeactions"

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeactions.Config, []string) error
}

func getCliCommands() map[string]cliCommand {
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
			description: "Show the next 20 location areas in the Pokemon world",
			callback:    CommandMap,
		},
		"mapBack": {
			name:        "mapBack",
			description: "Show the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <area-name>",
			description: "Show pokemons that live in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Try to catch the pokemon",
			callback:    commandCatch,
		},
	}
}
