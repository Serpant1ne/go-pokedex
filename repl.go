package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Serpant1ne/go-pokedex/internal/pokeactions"
	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
	"github.com/Serpant1ne/go-pokedex/internal/pokedex"
)

const (
	CACHE_INTERVAL = 30 * time.Second
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()
	config := pokeactions.Config{
		NextLocation:    "https://pokeapi.co/api/v2/location-area",
		PrevLocation:    "",
		BaseLocationUrl: "https://pokeapi.co/api/v2/location-area",
		BasePokemonUrl:  "https://pokeapi.co/api/v2/pokemon",
		Client: pokeactions.Client{
			Cache: pokecache.NewCache(CACHE_INTERVAL),
		},
		Pokedex: pokedex.NewPokedex(),
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		params := strings.Split(reader.Text(), " ")
		if len(params[0]) == 0 {
			continue
		}
		command, ok := cliCommands[params[0]]
		if !ok {
			commandNotFound()
			continue
		}
		err := command.callback(&config, params[1:])
		if err != nil {
			fmt.Println(err)
		}
	}
}
