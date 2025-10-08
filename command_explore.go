package main

import (
	"fmt"

	"github.com/Hircrown/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config) error {
	pokemonByLocation, err := pokeapi.PokemonList(cfg.pokemonLocation, cfg.cache)
	if err != nil {
		return err
	}

	for _, encounter := range pokemonByLocation.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
