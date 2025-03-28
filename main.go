package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next string
	Prev string
}

type areaLocation struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type areaLocationResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []areaLocation `json:"results"`
}

var commands = map[string]cliCommand{}

func init() {
	commands = map[string]cliCommand{
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
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	conf := config{
		Next: "https://pokeapi.co/api/v2/location-area",
		Prev: "",
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		raw := scanner.Text()
		clean := cleanInput(raw)
		if command, ok := commands[clean[0]]; ok {
			command.callback(&conf)
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

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, comm := range commands {
		fmt.Printf("%s: %s\n", comm.name, comm.description)
	}
	return nil
}

func commandMap(c *config) error {
	locations := areaLocationResponse{}
	res, err := http.Get(c.Next)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &locations); err != nil {
		return err
	}
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	c.Prev = c.Next
	c.Next = locations.Next
	return nil
}

func commandMapBack(c *config) error {
	if len(c.Prev) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}
	locations := areaLocationResponse{}
	res, err := http.Get(c.Prev)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &locations); err != nil {
		return err
	}
	for _, area := range locations.Results {
		fmt.Println(area.Name)
	}
	c.Prev = locations.Previous
	c.Next = locations.Next
	return nil
}
