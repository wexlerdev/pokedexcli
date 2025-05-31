package main

import (
	"time"
	"github.com/wexlerdev/pokedexcli/internal/pokeapi"
)




func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	
	config := config{
		pokeapiClient: pokeClient,
		pokedex: make(map[string]pokeapi.PokemonData),
	}
	startRepl(&config)
}



