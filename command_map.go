package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, params ...string) error {
    item, err := cfg.pokeapiClient.GetShallowLocation(cfg.nextLocationsURL)

    if err != nil {
        return err
    }

    cfg.nextLocationsURL = item.Next
    cfg.prevLocationsURL = item.Previous

    fmt.Println("Cities:")
    for _, result := range item.Results {
        fmt.Println("-", result.Name)
    }
    fmt.Println()

    return nil
}

func commandMapb(cfg *config, params ...string) error {
    if cfg.prevLocationsURL == nil {
        return errors.New("you're on the first page")
    }

    item, err := cfg.pokeapiClient.GetShallowLocation(cfg.prevLocationsURL)
    if err != nil {
        return err
    }

    cfg.nextLocationsURL = item.Next
    cfg.prevLocationsURL = item.Previous

    fmt.Println("Cities:")
    for _, result := range item.Results {
        fmt.Println("-", result.Name)
    }
    fmt.Println()

    return nil
}
