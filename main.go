package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jonackers/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
	caughtPokemon    map[string]pokeapi.PokemonResp
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 map locations",
			callback:    commandMapf,
		},
		"mapb": {
			name: 	  	 "mapb",
			description: "Lists the previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:		"explore",
			description: "Explore a location",
			callback:	commandExplore,
		},
		"catch": {
			name:		"catch",
			description: "Catch a Pokemon",
			callback:	commandCatch,
		},
		"inspect": {
			name:		"inspect",
			description: "Inspect a Pokemon you've caught",
			callback:	commandInspect,
		},
		"pokedex": {
			name:		"pokedex",
			description: "List all Pokemon you've caught",
			callback:	commandPokedex,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.PokemonResp),
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for a list of commands")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\nPokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}

		command, exists := getCommands()[input[0]]

		if exists {
			err := command.callback(cfg, input[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command...")
		}
	}
}
