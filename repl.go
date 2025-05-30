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




func startRepl(config *config) {
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

		err := cmd.callback(config)
		if err != nil {
			fmt.Printf("ERRORROROROR")
			os.Exit(1)
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
