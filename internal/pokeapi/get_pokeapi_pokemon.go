package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonURL string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonURL

	data, ok := c.cache.Get(url)
	if ok {
		var pokemonResp Pokemon
		if err := json.Unmarshal(data, &pokemonResp); err != nil {
			return Pokemon{}, err
		}
	
		return pokemonResp, nil
	}

    res, err := http.Get(url)
    if err != nil {
        return Pokemon{}, fmt.Errorf("error creating request: %w", err)
    }
    defer res.Body.Close()

    data, err = io.ReadAll(res.Body)
    if err != nil {
		return Pokemon{}, fmt.Errorf("error reading response: %w", err)
	}

	var pokemonResp Pokemon
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	return pokemonResp, nil
}