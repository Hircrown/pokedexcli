package main

import (
	"time"

	"github.com/Hircrown/pokedexcli/internal/pokecache"
)

const (
	cliCustomCmd  = "Pokedex > "
	cacheInterval = 30 * time.Second
)

func main() {
	cfg := &config{
		cache: pokecache.NewCache(cacheInterval),
	}
	startRepl(cliCustomCmd, cfg)
}
