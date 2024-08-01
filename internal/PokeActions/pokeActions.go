package pokeactions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
)

type locationAreaData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreaData(url string, c *pokecache.Cache) (locationAreaData, error) {
	locData := locationAreaData{}

	body := make([]byte, 0)

	if val, ok := c.Get(url); ok {
		body = val
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return locationAreaData{}, err
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			err := fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, body)
			return locationAreaData{}, err
		}
		if err != nil {
			return locationAreaData{}, err
		}

		c.Set(url, body)
	}

	err := json.Unmarshal(body, &locData)
	if err != nil {
		return locationAreaData{}, err
	}
	return locData, nil
}
