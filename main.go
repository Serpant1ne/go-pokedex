package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()
	config := config{
		next: "https://pokeapi.co/api/v2/location-area",
		prev: "",
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		command, ok := cliCommands[reader.Text()]
		if !ok {
			commandNotFound()
			continue
		}
		command.callback(&config)
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	next string
	prev string
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
			callback:    pokeactions.commandMap,
		},
		"mapBack": {
			name:        "mapBack",
			description: "Show the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
	}
}

func commandNotFound() error {
	fmt.Println("Error: wrong command")
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit(config *config) error {
	os.Exit(0)
	return nil

}
