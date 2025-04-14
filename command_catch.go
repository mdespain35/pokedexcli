package main

import (
	"fmt"
	"math/rand"
)

func commandCapture(cfg *config, pokemon string, pokedex *Pokedex) error {
	pokemonResp, err := cfg.pokeapiClient.LookUpPokemon(pokemon)
	if err != nil {
		return err
	}
	catchChance := 50.0 / pokemonResp.Experience
	fmt.Println("Throwing a Pokeball at " + pokemonResp.Name + "...")
	if rand.Float32() <= catchChance {
		fmt.Println(pokemonResp.Name + " was caught!")
		pokedex.caughtPokemon[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Println(pokemonResp.Name + " escaped!")
	}

	return nil
}
