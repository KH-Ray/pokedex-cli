package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"internal/pokeapi"
)

type config struct {
    pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
    pokedex          map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print("Pokedex > ")
        scanner.Scan()

        texts := cleanInput(scanner.Text())
        if len(texts) == 0 {
            continue
        }

        commandName := texts[0]
        cityName := ""

        if len(texts) > 1 {
            cityName = texts[1]
        }

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg, cityName)
            if err != nil {
                fmt.Println("Error:", err)
                fmt.Println()
            }
            continue
        } else {
            fmt.Println("Unknown command")
            fmt.Println()
			continue
        }
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "mapf": {
            name:        "mapf",
            description: "Display the names of the NEXT 20 location areas",
            callback:    commandMapf,
        },
        "mapb": {
            name:        "mapb",
            description: "Display the names of the PREVIOUS 20 location areas",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore",
            description: "Display a set of possible Pokémon encounters",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch",
            description: "Catch a Pokémon",
            callback:    commandCatch,
        },
        "inspect": {
            name:        "inspect",
            description: "Inspect a Pokémon",
            callback:    commandInpect,
        },
        "pokedex": {
            name:        "pokedex",
            description: "List all Pokémon that you have caught",
            callback:    commandPokedex,
        },
    }
}
