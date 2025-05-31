package main

import (
	"fmt"
	"os"
)

func commandExit(_ * config, _ ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ * config, _ ...string) error {
	helpString := `Welcome to the Pokedex!
Usage:
help: Displays a help message
exit: Exit the Pokedex` 
	fmt.Println(helpString)
	return nil
}

func commandMapf(config *config, params ...string) error {

	locationsData, err := config.pokeapiClient.GetLocationAreas(config.nextLocationsURL)
	if err != nil {
		return nil
	}

	config.nextLocationsURL = locationsData.Next
	config.prevLocationsURL = locationsData.Previous

	for _, loc := range locationsData.Results {
		fmt.Println(loc.Name)
	}
	return nil
}


func commandMapb(config *config, params ...string) error {
	locationsData, err := config.pokeapiClient.GetLocationAreas(config.prevLocationsURL)
	if err != nil {
		return nil
	}

	config.nextLocationsURL = locationsData.Next
	config.prevLocationsURL = locationsData.Previous

	for _, loc := range locationsData.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(config * config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("need to pass in location area name")
	}

	locationName := params[0]

	pokemonNameSlice, err := config.pokeapiClient.GetPokemonInArea(locationName)
	if err != nil {
		return err
	}

	for _, pokemonName := range pokemonNameSlice {
		fmt.Println(pokemonName)
	}
	return nil
}
