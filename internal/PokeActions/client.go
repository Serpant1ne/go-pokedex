package pokeactions

import (
	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
	"github.com/Serpant1ne/go-pokedex/internal/pokedex"
)

type Client struct {
	Cache pokecache.Cache
}

type Config struct {
	Client          Client
	Pokedex         pokedex.Pokedex
	NextLocation    string
	PrevLocation    string
	BaseLocationUrl string
	BasePokemonUrl  string
}
