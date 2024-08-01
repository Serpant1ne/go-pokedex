package main

import "github.com/Serpant1ne/go-pokedex/internal/pokeactions"

func main() {
	startRepl()
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	client pokeactions.Client
	next   string
	prev   string
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
	}
}
