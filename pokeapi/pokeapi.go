package pokeapi

import (
	"http"
	"encoding/json"
	"fmt"
)

type PokeLocationData struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

const pokeBaseApiUrl = "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"

func getLocationAreas(url string)([]string, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Location GET failed")
	}
	
	var pokeLocationStruct PokeLocationData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeLocationStruct)
	if err != nil {
		return nil, fmt.Errof("Failure Decoding Location Bytes")
	}

	locations := make([]string, 0,20)
	for i, resultStruct := range pokeLocationStruct.Results {
		locations := append(locations,resultsStruct.Name)
	}
	return locations, nil
}
