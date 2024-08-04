package pokeactions

import "sync"

type Pokedex struct {
	Pokemons map[string]pokemon
	Mux      *sync.Mutex
}

func NewPokedex() Pokedex {
	return Pokedex{
		Pokemons: make(map[string]pokemon),
		Mux:      &sync.Mutex{},
	}
}
