package main

import (
	"fmt"
	"os"
	"math/rand/v2"
	"github.com/wexlerdev/pokedexcli/internal/pokeapi"
	"strings"
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
	if !didCatch {
		fmt.Println("Did not Catch :(😫")
		return nil
	} 

	//did catch :-)
	config.pokedex[pokemonName] = *pokemonData
	fmt.Println("Caught and added to pokedex 😁")


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

func commandInspect(config * config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("need to pass in pokemon name")
	}

	pokemonName := params[0]

	pokemon, found := config.pokedex[pokemonName]
	if !found {
		fmt.Println("you have not caught this pokemon silly")
		return nil
	}
	printPokemonDetails(pokemon)

	return nil
}

func printPokemonDetails(p pokeapi.PokemonData) {
	fmt.Println("--- Pokémon Stats ---")
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Height: %d dm (%s feet, %.1f inches)\n", p.Height, formatHeight(p.Height), float64(p.Height)*3.937)
	fmt.Printf("Weight: %d hg (%s lbs)\n", p.Weight, formatWeight(p.Weight))
	fmt.Printf("Front Default Sprite: %s\n", p.Sprites.FrontDefault)

	fmt.Println("\nBase Stats:")
	for _, stat := range p.Stats {
		fmt.Printf("  %s: %d\n", formatStatName(stat.Stat.Name), stat.BaseStat)
	}
}

// formatHeight converts decimeters to feet and inches.
func formatHeight(decimeters int) string {
	totalInches := float64(decimeters) * 3.937
	feet := int(totalInches / 12)
	inches := int(totalInches) % 12
	return fmt.Sprintf("%d'%d\"", feet, inches)
}

// formatWeight converts hectograms to pounds.
func formatWeight(hectograms int) string {
	pounds := float64(hectograms) * 0.220462
	return fmt.Sprintf("%.1f lbs", pounds)
}

// formatStatName formats stat names like "special-attack" to "Special Attack".
func formatStatName(s string) string {
	s = strings.ReplaceAll(s, "-", " ") // Replace hyphens with spaces
	return strings.Title(s)             // Capitalize first letter of each word
}
