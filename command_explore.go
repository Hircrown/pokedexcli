package main

import (
	"fmt"

	"github.com/Hircrown/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config) error {
	pokemons, err := pokeapi.PokemonList(cfg.pokemonLocation, cfg.cache)
	if err != nil {
		return err
	}

	for _, pokemon := range pokemons.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
