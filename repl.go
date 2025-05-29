package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"github.com/wexlerdev/pokedexcli/internal/pokeapi"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}



func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	stringSlice := make([]string, 0)
	stringSlice = strings.Fields(lowerText)
	return stringSlice
}




func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex > ")
		scanner.Scan() 
		line := scanner.Text()
		stringSlice := cleanInput(line)
		
		if stringSlice == nil {
			fmt.Println("stringslice nil ")
			continue
		}

		if len(stringSlice) == 0 {
			fmt.Println("stringslice empty ")
			continue
		}

		cmd, ok := getCommands()[stringSlice[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.callback(&config)
		if err != nil {
			fmt.Printf("ERRORROROROR")
			os.Exit(1)
		}
	}
}

