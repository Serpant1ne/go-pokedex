package pokeactions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Serpant1ne/go-pokedex/internal/pokecache"
)

type location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetLocation(url string, name string, c *pokecache.Cache) (location, error) {
	locData := location{}

	body := make([]byte, 0)
	link := fmt.Sprintf("%s/%s/", url, name)

	if val, ok := c.Get(link); ok {
		body = val
	} else {
		resp, err := http.Get(link)
		if err != nil {
			return location{}, err
		}

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			err := fmt.Errorf("response failed with status code: %d and\nbody: %s", resp.StatusCode, body)
			return location{}, err
		}
		if err != nil {
			return location{}, err
		}

		c.Set(url, body)
	}

	err := json.Unmarshal(body, &locData)
	if err != nil {
		return location{}, err
	}
	return locData, nil
}
