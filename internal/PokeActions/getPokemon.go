package pokeactions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
	"github.com/Serpant1ne/go-pokedex/internal/pokedex"
)

func GetPokemon(url string, name string, c *pokecache.Cache) (pokedex.Pokemon, error) {
	pokemonData := pokedex.Pokemon{}

	body := make([]byte, 0)
	link := fmt.Sprintf("%s/%s/", url, name)

	if val, ok := c.Get(link); ok {
		body = val
	} else {
		resp, err := http.Get(link)
		if err != nil {
			return pokedex.Pokemon{}, err
		}

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			err := fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, body)
			return pokedex.Pokemon{}, err
		}
		if err != nil {
			return pokedex.Pokemon{}, err
		}

		c.Set(url, body)
	}

	err := json.Unmarshal(body, &pokemonData)
	if err != nil {
		return pokedex.Pokemon{}, err
	}
	return pokemonData, nil
}
