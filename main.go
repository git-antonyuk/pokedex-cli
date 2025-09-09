package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/git-antonyuk/pokedex-cli/internal"
	api_get_location_area_details "github.com/git-antonyuk/pokedex-cli/internal/api/get_location_area_details"
	api_get_location_areas "github.com/git-antonyuk/pokedex-cli/internal/api/get_location_areas"
	string_utils_print_list "github.com/git-antonyuk/pokedex-cli/internal/string_utils"
)

const COMMAND_LINE_NAME = "Pokedex > "

type configCommand struct {
	cliCommandsMap 	map[string]cliCommand
	Next           	*string
	Previous       	*string
	cache 			pokecache.Cache
	fullCommand 	[]string
}
type cliCommand struct {
	name        	string
	description 	string
	callback    	func(*configCommand) error
}

func main() {
	fmt.Print(COMMAND_LINE_NAME)
	scanner := bufio.NewScanner(os.Stdin)
	cache := pokecache.NewCache(10 * time.Minute)

	commandsMap := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Go back to prev location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore details of location areas, use name param",
			callback:    commandExplore,
		},
	}

	next := ""
	previous := ""
	config := configCommand{
		cliCommandsMap: commandsMap,
		Next:           &next,
		Previous:       &previous,
		cache: 			cache,
	}

	for scanner.Scan() {
		userText := scanner.Text()
		slicedText := cleanInput(userText)
		possibleCommand := slicedText[0]
		command, ok := commandsMap[possibleCommand]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			config.fullCommand = slicedText
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Print(COMMAND_LINE_NAME)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func commandExit(config *configCommand) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// !!!TODO move to utils
func getCommandsDescriptions(cliCommands map[string]cliCommand) string {
	var sb strings.Builder
	for _, cmd := range cliCommands {
		fmt.Fprintf(&sb, "%s: %s\n", cmd.name, cmd.description)
	}
	return sb.String()
}

func commandHelp(config *configCommand) error {
	var sb strings.Builder
	fmt.Fprint(&sb, "Welcome to the Pokedex!\n")
	fmt.Fprint(&sb, "Usage:\n\n")
	sb.WriteString(getCommandsDescriptions(config.cliCommandsMap))
	res := sb.String()
	fmt.Print(res)

	return nil
}

func cleanInput(text string) []string {
	lowerCased := strings.ToLower(text)
	return strings.Fields(lowerCased)
}

// TODO: move to utils
func getLocationsAndPrintResults(url string, config *configCommand) error {
	locationData, _ := api_get_location_areas.GetLocationAreas(config.cache, url)
	// Update Next and Previous
	*config.Next = locationData.Next
	*config.Previous = locationData.Previous
	// Convert list and print it
	locationsList := api_get_location_areas.ConvertLocationToNameList(locationData.Results)
	string_utils_print_list.PrintList(locationsList)
	return nil
}

func commandMap(config *configCommand) error {
	return getLocationsAndPrintResults(*config.Next, config)
}

func commandMapBack(config *configCommand) error {
	if *config.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return getLocationsAndPrintResults(*config.Previous, config)
}

func commandExplore(config *configCommand) error {
	if len(config.fullCommand) < 2 {
		return errors.New("You forgot to add name as second parameter")
	}
	name := config.fullCommand[1]
	fmt.Printf("Exploring %v...\n", name)
	locationDetails, _ := api_get_location_area_details.GetLocationAreaDetails(config.cache, name)
	pokemonList := api_get_location_area_details.GetPokemonsList(locationDetails)
	string_utils_print_list.PrintList(pokemonList)
	return nil
}
