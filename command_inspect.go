package main

import (
	"errors"
	"fmt"
)

func commandInpect(cfg *config, params ...string) error {
    pokemonName := params[0]

    if len(pokemonName) == 0 {
        return errors.New("pokemon name is empty")
    }

    val, ok := cfg.pokedex[pokemonName]

    if !ok {
        fmt.Println("You have not caught that pokemon")
    } else {
        fmt.Printf("Name: %s\n", val.Name)
        fmt.Printf("Height: %d\n", val.Height)
        fmt.Printf("Weight: %d\n", val.Weight)

        fmt.Println("Stats:")
        for _, pokemonStats := range val.Stats {
            fmt.Printf("  - %s: %d\n", pokemonStats.Stat.Name, pokemonStats.BaseStat)
        }

        fmt.Println("Types:")
        for _, pokemonType := range val.Types {
            fmt.Println("  -", pokemonType.Type.Name)
        }
    }
    fmt.Println()

    return nil
}

