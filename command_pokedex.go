package main

import (
	"fmt"
)

func commandPokedex(cfg *config) error {
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("you pokedex is still empty")
	}
	fmt.Println("Your Pokedex:")
	for pokemon, _ := range cfg.pokedex {
		fmt.Println(" -", pokemon)
	}
	return nil
}
