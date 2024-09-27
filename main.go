package main

import (
	"time"

	"internal/pokeapi"
)

func main() {
    pokeClient := pokeapi.NewClient(time.Hour)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex: make(map[string]pokeapi.Pokemon),
	}

    startRepl(cfg)
}
