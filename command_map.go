package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsResp.Next
	cfg.prevLocationURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
