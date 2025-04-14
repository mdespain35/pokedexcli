package main

import (
	"fmt"
	"strconv"
)

func commandInspect(cfg *config, pokemon string, pokedex *Pokedex) error {
	poke, err := pokedex.caughtPokemon[pokemon]
	if !err {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println(
		"Name: " + poke.Name +
			"\nHeight: " + strconv.Itoa(poke.Height) +
			"\nWeight: " + strconv.Itoa(poke.Weight) +
			"\nStats:")
	for _, s := range poke.Stats {
		fmt.Println(" -" + s.Stat.Name + ": " + strconv.Itoa(s.Base))
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Println(" - " + t.Type.Name)
	}

	return nil
}
