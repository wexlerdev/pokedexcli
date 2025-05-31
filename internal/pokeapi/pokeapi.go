package pokeapi

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
)


const (
	baseURL = "https://pokeapi.co/api/v2"
)


func (c *Client) GetLocationAreas(pageUrl *string)(PokeLocationData, error) {

	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	//check cache for url
	byteSlice, isCacheHit := c.cache.Get(url)

	if isCacheHit {
		fmt.Println("CACHE HIT")
		var pokeLocationStruct PokeLocationData
		err := json.Unmarshal(byteSlice, &pokeLocationStruct)
		if err != nil {
			return PokeLocationData{}, fmt.Errorf("error unmarshaling cache hit")
		}
		return pokeLocationStruct, nil
	}

	//Cache Miss ------
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLocationData{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLocationData{}, fmt.Errorf("Location GET failed")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PokeLocationData{}, err
	}

	var pokeLocationStruct PokeLocationData

	err = json.Unmarshal(data, &pokeLocationStruct)
	if err != nil {
		return PokeLocationData{}, fmt.Errorf("Failure Decoding Location Bytes")
	}

	//store data in cache
	c.cache.Add(url, data)

	return pokeLocationStruct, nil
}


func (c * Client) GetPokemonInArea(locationName string) ([]string, error) {
	url := baseURL + "/location-area/" + locationName
	
	//check cache for url
	byteSlice, isCacheHit := c.cache.Get(url)

	if isCacheHit {
		fmt.Println("CACHE HIT")
		var pokeEncounterStruct PokemonEncounterData
		err := json.Unmarshal(byteSlice, &pokeEncounterStruct)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling cache hit")
		}

		var pokemonNameSlice []string
		for _, pokemonStruct := range pokeEncounterStruct.PokemonEncounters {
			pokemonNameSlice = append(pokemonNameSlice, pokemonStruct.Pokemon.Name)
		}
		return pokemonNameSlice, nil
	}

	//Cache Miss ------
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pokeEncounterStruct PokemonEncounterData

	err = json.Unmarshal(data, &pokeEncounterStruct)
	if err != nil {
		return nil, err
	}


	var pokemonNameSlice []string
	for _, pokemon := range pokeEncounterStruct.PokemonEncounters {
		pokemonNameSlice = append(pokemonNameSlice, pokemon.Pokemon.Name)
	}

	//add to cache
	c.cache.Add(url, data)

	return pokemonNameSlice, nil
}

func (c * Client) GetPokemon(pokemonName string) (* PokemonData, error) {
	
	url := baseURL + "/pokemon/" + pokemonName
	

	//Cache Miss ------
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pokemonData PokemonData

	err = json.Unmarshal(data, &pokemonData)
	if err != nil {
		return nil, err
	}

	return &pokemonData, nil
}
