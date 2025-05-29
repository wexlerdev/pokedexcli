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

const (
	baseURL = "https://pokeapi.co/api/v2"
)


func (c *Client) GetLocationAreas(pageUrl *string)(PokeLocationData, error) {

	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLocationData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLocationData{}, fmt.Errorf("Location GET failed")
	}
	defer res.Body.Close()
	
	var pokeLocationStruct PokeLocationData

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&pokeLocationStruct)
	if err != nil {
		return PokeLocationData{}, fmt.Errorf("Failure Decoding Location Bytes")
	}

	return pokeLocationStruct, nil
}
