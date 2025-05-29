package main 
import (
	"strings"
	"fmt"
	"bufio"
	"os"
	"github.com/wexlerdev/pokedexcli/pokeapi"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	stringSlice := make([]string, 0)
	stringSlice = strings.Fields(lowerText)
	return stringSlice
}

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}

type config struct {
	MapPage int
}


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

func mapWithNoPageChange(config *config) error {
	const pokeBaseApiUrl = "https://pokeapi.co/api/v2/location-area?limit=20"
	offset := (config.MapPage - 1) * 20

	pokeLocationApiUrl := pokeBaseApiUrl + fmt.Sprintf("&offset=%v", offset)

	locations, err := pokeapi.GetLocationAreas(pokeLocationApiUrl)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location)
	}
	
	return nil
}

func commandMap(config *config) error {
	//increment page
	config.MapPage++

	err := mapWithNoPageChange(config)
	if err != nil {
		return err
	}

	
	return nil
}

func commandBmap(config *config) error {
	if config.MapPage == 0 {
		config.MapPage = 1
	}

	if config.MapPage > 1 {
		config.MapPage--
	}

	err := mapWithNoPageChange(config)
	if err != nil {
		return err
	}

	return nil
}


func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := config{0}

	commandRegistry := map[string]cliCommand {
		"exit": {
			name:		"exit",
			description: "Exit the Pokedex",
			callback:	 commandExit,
		},
		"help": {
			name:		"help",
			description:	"helps with cli tool",
			callback: commandHelp,
		},
		"map": {
			name:			"map",
			description:	"shows the next 20 locations",
			callback:		commandMap,
		},
		"bmap": {
			name:			"bmap",
			description:	"shows the last 20 locations",
			callback:		commandBmap,
		},
	}

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

		cmd, ok := commandRegistry[stringSlice[0]]
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

