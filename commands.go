package main

import (
	"fmt"
	"os"

	"github.com/Serpant1ne/go-pokedex/internal/pokeactions"
)

func commandNotFound() error {
	fmt.Println("Error: wrong command")
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCliCommands() {
		fmt.Printf("%s: %s \n", command.name, command.description)
	}

	return nil
}

func commandExit(config *config) error {
	os.Exit(0)
	return nil
}

func CommandMap(config *config) error {
	locData, err := pokeactions.GetLocationAreaData(config.next)
	if err != nil {
		return err
	}
	for _, location := range locData.Results {
		fmt.Println(location.Name)
	}
	config.next = locData.Next
	config.prev = locData.Previous
	return nil
}

func commandMapBack(config *config) error {
	locData, err := pokeactions.GetLocationAreaData(config.prev)
	if err != nil {
		return err
	}
	for _, location := range locData.Results {
		fmt.Println(location.Name)
	}
	config.next = locData.Next
	config.prev = locData.Previous
	return nil
}
