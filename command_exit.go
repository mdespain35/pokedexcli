package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, exit string, pokedex *Pokedex) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
