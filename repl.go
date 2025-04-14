package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mdespain35/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationURL *string
	prevLocationURL *string
}

type Pokedex struct {
	caughtPokemon map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	pokedex := Pokedex{
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandArg := ""
		commandName := words[0]
		if len(words) > 1 {
			commandArg = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, commandArg, &pokedex)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	rawText := strings.Split(strings.TrimSpace(text), " ")
	cleaned := []string{}
	for _, r := range rawText {
		if r != "" {
			cleaned = append(cleaned, strings.ToLower(r))
		}
	}
	return cleaned
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string, *Pokedex) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
			description: "Displays the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Displays the pokemon that can be found at a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the pokemon you name",
			callback:    commandCapture,
		},
	}
}
