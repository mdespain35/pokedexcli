package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		raw := scanner.Text()
		clean := cleanInput(raw)
		if command, ok := commands[clean[0]]; ok {
			command.callback()
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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, comm := range commands {
		fmt.Printf("%s: %s\n", comm.name, comm.description)
	}
	return nil
}
