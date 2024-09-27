package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, params ...string) error {
    pokemonName := params[0]

    if len(pokemonName) == 0 {
        return errors.New("pokemon name is empty")
    }

    item, err := cfg.pokeapiClient.GetPokemon(pokemonName)

    if err != nil {
        return err
    }

    const threshold = 50
    randNum := rand.Intn(item.BaseExperience)

    fmt.Printf("Throwing a Pokeball at %s...\n", item.Name)
    time.Sleep(time.Second * 2)
    if randNum > threshold {
        fmt.Printf("%s escaped!\n", item.Name)
    } else {
        fmt.Printf("%s was caught!\n", item.Name)
        cfg.pokedex[pokemonName] = item
    }
    fmt.Println()

    return nil
}
