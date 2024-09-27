package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params ...string) error {
    cityName := params[0]

    if len(cityName) == 0 {
        return errors.New("city name is empty")
    }

    item, err := cfg.pokeapiClient.GetDeepLocation(cityName)

    if err != nil {
        return err
    }

    fmt.Println("Found Pokemon:")
    for _, pokemon := range item.PokemonEncounters {
        fmt.Println("-", pokemon.Pokemon.Name)
    }
    fmt.Println()

    return nil
}
