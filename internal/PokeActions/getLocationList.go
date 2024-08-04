package pokeactions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
)

type locationList struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationList(url string, c *pokecache.Cache) (locationList, error) {
	locData := locationList{}

	body := make([]byte, 0)

	if val, ok := c.Get(url); ok {
		body = val
	} else {
		resp, err := http.Get(url)
		if err != nil {
			return locationList{}, err
		}

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			err := fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, body)
			return locationList{}, err
		}
		if err != nil {
			return locationList{}, err
		}

		c.Set(url, body)
	}

	err := json.Unmarshal(body, &locData)
	if err != nil {
		return locationList{}, err
	}
	return locData, nil
}
