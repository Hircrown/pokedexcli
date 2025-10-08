package main

import (
	"time"

	"github.com/Hircrown/pokedexcli/internal/pokeapi"
	"github.com/Hircrown/pokedexcli/internal/pokecache"
)

const (
	cliCustomCmd  = "Pokedex > "
	cacheInterval = 30 * time.Second
)

func main() {
	cfg := &config{
		cache:   pokecache.NewCache(cacheInterval),
		pokedex: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cliCustomCmd, cfg)
}
