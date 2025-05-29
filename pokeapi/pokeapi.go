package pokeapi

import (
	"net/http"
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


func GetLocationAreas(url string)([]string, error) {

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Location GET failed")
	}
	
	var pokeLocationStruct PokeLocationData
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeLocationStruct)
	if err != nil {
		return nil, fmt.Errorf("Failure Decoding Location Bytes")
	}

	locations := make([]string, 0,20)
	for _, resultsStruct := range pokeLocationStruct.Results {
		locations = append(locations,resultsStruct.Name)
	}
	return locations, nil
}
