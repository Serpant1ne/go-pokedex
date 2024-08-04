package main

import (
	"errors"
	"fmt"
	"math/rand"
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
	locData, err := pokeactions.GetLocationList(config.NextLocation, &config.Client.Cache)
	if err != nil {
		return err
	}
	for _, location := range locData.Results {
		fmt.Println(location.Name)
	}
	config.NextLocation = locData.Next
	config.PrevLocation = locData.Previous
	return nil
}

func commandMapBack(config *pokeactions.Config, args []string) error {
	if config.PrevLocation == "" {
		return errors.New("error. you are on the first page")
	}
	locData, err := pokeactions.GetLocationList(config.PrevLocation, &config.Client.Cache)
	if err != nil {
		return err
	}
	for _, location := range locData.Results {
		fmt.Println(location.Name)
	}
	config.NextLocation = locData.Next
	config.PrevLocation = locData.Previous
	return nil
}

func commandExplore(config *pokeactions.Config, args []string) error {
	if config.BaseLocationUrl == "" {
		return errors.New("error. No BaseUrl")
	}
	locData, err := pokeactions.GetLocation(config.BaseLocationUrl, args[0], &config.Client.Cache)
	if err != nil {
		return err
	}
	for _, pokEncounter := range locData.PokemonEncounters {
		fmt.Println(pokEncounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(config *pokeactions.Config, args []string) error {
	if config.BasePokemonUrl == "" {
		return errors.New("error. No BaseUrl")
	}
	pokemon, err := pokeactions.GetPokemon(config.BasePokemonUrl, args[0], &config.Client.Cache)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball to %s...\n", pokemon.Name)
	if rand.Intn(pokemon.BaseExperience) < 50 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.Pokedex.AddPokemon(pokemon)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func commandInspect(config *pokeactions.Config, args []string) error {
	pokemon, ok := config.Pokedex.GetPokemon(args[0])
	if !ok {
		return errors.New("error. You don't have this pokemon")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}
