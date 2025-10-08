package main

import (
	"fmt"
	"math/rand"

	"github.com/Hircrown/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config) error {
	pokemon, err := pokeapi.GetPokemonStats(cfg.pokemonName, cfg.cache)
	if err != nil {
		return err
	}

	res := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if res > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}
