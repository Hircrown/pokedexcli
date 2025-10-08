package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Hircrown/pokedexcli/internal/pokecache"
)

func GetPokemonStats(pokemonName string, cache pokecache.Cache) (Pokemon, error) {
	var pokemon Pokemon
	fullURL := fmt.Sprintf("%spokemon/%s", baseURL, pokemonName)

	if pokemonStats, exists := cache.Get(fullURL); exists {
		if err := json.Unmarshal(pokemonStats, &pokemon); err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshaling pokemon: %w", err)
		}
		return pokemon, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error requesting pokemon: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading pokemon: %w", err)
	}
	if err = json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling pokemon: %w", err)
	}
	cache.Add(fullURL, data)

	return pokemon, nil
}
