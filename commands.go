package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Serpant1ne/go-pokedex/internal/pokeactions"
)

func commandNotFound() error {
	fmt.Println("Error: wrong command")
	return nil
}

func commandHelp(config *pokeactions.Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCliCommands() {
		fmt.Printf("%s: %s \n", command.name, command.description)
	}

	return nil
}

func commandExit(config *pokeactions.Config, args []string) error {
	os.Exit(0)
	return nil
}

func CommandMap(config *pokeactions.Config, args []string) error {
	locData, err := pokeactions.GetLocationList(config.Next, &config.Client.Cache)
	if err != nil {
		return err
	}
	for _, location := range locData.Results {
		fmt.Println(location.Name)
	}
	config.Next = locData.Next
	config.Prev = locData.Previous
	return nil
}

func commandMapBack(config *pokeactions.Config, args []string) error {
	if config.Prev == "" {
		return errors.New("error. you are on the first page")
	}
	locData, err := pokeactions.GetLocationList(config.Prev, &config.Client.Cache)
	if err != nil {
		return err
	}
	for _, location := range locData.Results {
		fmt.Println(location.Name)
	}
	config.Next = locData.Next
	config.Prev = locData.Previous
	return nil
}

func commandExplore(config *pokeactions.Config, args []string) error {
	if config.BaseUrl == "" {
		return errors.New("error. No BaseUrl")
	}
	locData, err := pokeactions.GetLocation(config.BaseUrl, args[0], &config.Client.Cache)
	if err != nil {
		return err
	}
	for _, pokEncounter := range locData.PokemonEncounters {
		fmt.Println(pokEncounter.Pokemon.Name)
	}
	return nil
}
