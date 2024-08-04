package pokeactions

import "github.com/Serpant1ne/go-pokedex/internal/pokecache"

type Client struct {
	Cache pokecache.Cache
}

type Config struct {
	Client  Client
	Next    string
	Prev    string
	BaseUrl string
}
