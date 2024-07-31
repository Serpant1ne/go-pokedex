package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
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
