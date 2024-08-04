package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Serpant1ne/go-pokedex/internal/pokeactions"
	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
)

const (
	CACHE_INTERVAL = 30 * time.Second
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	cliCommands := getCliCommands()
	config := pokeactions.Config{
		Next: "https://pokeapi.co/api/v2/location-area",
		Prev: "",
		Client: pokeactions.Client{
			Cache: pokecache.NewCache(CACHE_INTERVAL),
		},
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
