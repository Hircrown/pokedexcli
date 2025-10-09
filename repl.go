package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Hircrown/pokedexcli/internal/pokeapi"
	"github.com/Hircrown/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	lower := strings.ToLower(text)
	return strings.Fields(lower)
}

func startRepl(placeholder string, cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(placeholder)
		scanner.Scan()

		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		cmdName := words[0]

		if cmdName == "explore" || cmdName == "inspect" {
			if len(words) != 2 {
				fmt.Printf("Usage: %s location-area -> Example: %s canalave-city-area\n", cmdName, cmdName)
			} else {
				cfg.pokemonLocation = words[1]
			}
		}

		if cmdName == "catch" {
			if len(words) != 2 {
				fmt.Printf("Usage: %s pokemonName -> Example: %s mewtwo\n", cmdName, cmdName)
				continue
			} else {
				cfg.pokemonName = words[1]
			}
		}

		cmd, exists := getCommands()[cmdName]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			if err := cmd.callback(cfg); err != nil {
				fmt.Printf("Callback error: %v\n", err)
			}
		}
	}
}

type config struct {
	cache           pokecache.Cache
	previous        *string
	next            *string
	pokemonLocation string
	pokemonName     string
	pokedex         map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Display the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all the Pokemon located in a specific area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show the pokemon info you have in your Pokedex",
			callback:    commandInspect,
		},
		"help": {
			name:        "help",
			description: "Give an overview of cli commands",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
