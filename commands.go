package main

import (
	"fmt"
	"os"
	"math/rand/v2"
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

func commandCatch(config * config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("need to pass in pokemon name")
	}

	pokemonName := params[0]

	pokemonData, err := config.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	//calculate chance
	//assume an XP of 0 means 100 percent chance and an XP of 1000 is a 10 percent chance
	chanceToCatch := calculateCatchChance(pokemonData.BaseExperience)
	didCatch := rand.Float64() - chanceToCatch <= 0.0
	fmt.Printf("Throwing a Pokeball at %v...", pokemonName)
	fmt.Printf("Chance to catch %v: %v%%\n", pokemonName, chanceToCatch * 100.0)
	if didCatch {
		fmt.Println("CAUGHT!!")
	} else {
		fmt.Println("DID NOT CATCH")
	}

	return nil

}

func calculateCatchChance(xp int) float64 {
	//xp of 0 is 100 percent chance to catch
	//xp of 1000 is 10 percent chance to catch
	//these are two points (0,100) (1000,10)
	//where x is the xp and y is the chance to catch
	// y = -0.09 * x + 100.0

	slope := -0.09
	chanceToCatch := slope * float64(xp) + 100.0

	if chanceToCatch > 100.0 {
		return 1.0
	}

	if chanceToCatch < 10.0 {
		return 0.1
	}
	return chanceToCatch / 100.0
}
