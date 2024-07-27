package pokeactions

import "net/http"

type locationAreaData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *config) error {
	resp, err := http.Get(config.next)
	if err != nil {
		return err
	}
}
