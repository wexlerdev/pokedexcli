package main

import (
	"fmt"
	"os"
)

func commandExit(_ * config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ * config) error {
	helpString := `Welcome to the Pokedex!
Usage:
help: Displays a help message
exit: Exit the Pokedex` 
	fmt.Println(helpString)
	return nil
}

func commandMapf(config *config) error {

	locationsData, err := config.pokeapiClient.GetLocationAreas(config.nextLocationsURL)
	if err != nil {
		return nil
	}

	config.nextLocationsURL = locationsData.Next
	config.prevLocationsURL = locationsData.Previous

	for _, loc := range locationsData.Results {
		fmt.Println(loc.Name)
	}
}


func commandMapb(config *config) error {
	locationsData, err := config.pokeapiClient.GetLocationAreas(config.prevLocationsURL)
	if err != nil {
		return nil
	}

	config.nextLocationsURL = locationsData.Next
	config.prevLocationsURL = locationsData.Previous

	for _, loc := range locationsData.Results {
		fmt.Println(loc.Name)
	}
}
