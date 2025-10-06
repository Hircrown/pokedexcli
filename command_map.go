package main

import (
	"fmt"

	"github.com/Hircrown/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	locations, err := pokeapi.LocationAreaList(cfg.next)
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
