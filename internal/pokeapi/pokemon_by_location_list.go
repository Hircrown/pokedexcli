package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Hircrown/pokedexcli/internal/pokecache"
)

func PokemonList(location string, cache pokecache.Cache) (PokemonByLocation, error) {
	fullURL := fmt.Sprintf("%slocation-area/%s", baseURL, location)
	var pokemonByLocation PokemonByLocation

	if data, exists := cache.Get(fullURL); exists {
		if err := json.Unmarshal(data, &pokemonByLocation); err != nil {
			return PokemonByLocation{}, fmt.Errorf("error unmarshaling pokemon: %w", err)
		}
		return pokemonByLocation, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return PokemonByLocation{}, fmt.Errorf("error requesting pokemon: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonByLocation{}, fmt.Errorf("error reading pokemon response: %w", err)
	}

	if err := json.Unmarshal(data, &pokemonByLocation); err != nil {
		return PokemonByLocation{}, fmt.Errorf("error unmarshaling pokemon: %w", err)
	}

	cache.Add(fullURL, data)

	return pokemonByLocation, nil
}
