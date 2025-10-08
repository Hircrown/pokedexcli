package main

import (
	"fmt"

	"github.com/Hircrown/pokedexcli/internal/pokeapi"
)

func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		return fmt.Errorf("you are on the first page")
	}
	locations, err := pokeapi.LocationAreaList(cfg.previous, cfg.cache)
	if err != nil {
		return err
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	cfg.previous = locations.Previous
	cfg.next = locations.Next

	return nil
}
