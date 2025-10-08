package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Hircrown/pokedexcli/internal/pokecache"
)

func LocationAreaList(pageURL *string, cache pokecache.Cache) (LocationArea, error) {
	fullURL := baseURL + "location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	var locations LocationArea
	//cache usage logic
	if data, exists := cache.Get(fullURL); exists {
		fmt.Println("\nusing cache")
		if err := json.Unmarshal(data, &locations); err != nil {
			return LocationArea{}, fmt.Errorf("error unmarshaling location data: %w", err)
		}
		return locations, nil
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error requesting location area: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading location area response: %w", err)
	}
	if err = json.Unmarshal(data, &locations); err != nil {
		return LocationArea{}, fmt.Errorf("error unmarshaling location data: %w", err)
	}
	cache.Add(fullURL, data)

	return locations, nil
}
