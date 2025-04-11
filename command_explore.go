package main

import (
	"fmt"
)

func commandExplore(cfg *config, explore string) error {
	pokemonResp, err := cfg.pokeapiClient.ListPokemon(explore)
	if err != nil {
		return err
	}
	fmt.Println("Exploring " + explore + "...")
	if len(pokemonResp.Encounters) == 0 {
		fmt.Println("No Pokemon found!")
	} else {
		fmt.Println("Found Pokemon:")
	}
	for _, pokemon := range pokemonResp.Encounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}
