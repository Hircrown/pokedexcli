package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LocationAreaList(pageURL *string) (LocationArea, error) {
	fullURL := baseURL + "location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	res, err := http.Get(fullURL)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error requesting location area: %w", err)
	}
	defer res.Body.Close()

	var locations LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&locations); err != nil {
		return LocationArea{}, fmt.Errorf("error decoding location area response: %w", err)
	}

	return locations, nil
}
