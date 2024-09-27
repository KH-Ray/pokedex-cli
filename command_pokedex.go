package main

import "fmt"

func commandPokedex(cfg *config, params ...string) error {
    fmt.Println("Your Pokedex:")
    
    for _, v := range cfg.pokedex {
        fmt.Printf("  - %s\n", v.Name)
    }
    fmt.Println()

    return nil
}
