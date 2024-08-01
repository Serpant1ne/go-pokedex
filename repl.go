package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Serpant1ne/go-pokedex/internal/pokeactions"
	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
)

const (
	CACHE_INTERVAL = 5 * time.Second
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()
	config := config{
		next: "https://pokeapi.co/api/v2/location-area",
		prev: "",
		client: pokeactions.Client{
			Cache: pokecache.NewCache(CACHE_INTERVAL),
		},
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		if len(reader.Text()) == 0 {
			continue
		}
		command, ok := cliCommands[reader.Text()]
		if !ok {
			commandNotFound()
			continue
		}
		command.callback(&config)
	}
}
