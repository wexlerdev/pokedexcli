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
	callback	func() error
}

type config struct {
	MapPage int
}


func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
func commandHelp() error {
	helpString := `Welcome to the Pokedex!
Usage:
help: Displays a help message
exit: Exit the Pokedex` 
	fmt.Println(helpString)
	return nil
}

func commandMap(config *config) error {
	const pokeBaseApiUrl = "https://pokeapi.co/api/v2/location-area?limit=20"
	offset := (config.MapPage - 1) * 20

	pokeLocationApiUrl := pokeBaseApiUrl + fmt.Sprintf("&offset=%v", offset)


}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		err := cmd.callback()
		if err != nil {
			fmt.Printf("ERRORROROROR")
			os.Exit(1)
		}
	}
}

