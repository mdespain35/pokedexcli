package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemon(locationName string) (RespShallowPokemonEnc, error) {
	if len(locationName) == 0 {
		return RespShallowPokemonEnc{}, errors.New("please enter a location name to explore")
	}
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemonEnc{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemonEnc{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemonEnc{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemonEnc{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemonEnc{}, err
	}

	pokemonResp := RespShallowPokemonEnc{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return RespShallowPokemonEnc{}, err
	}

	c.cache.Add(url, data)
	return pokemonResp, nil
}
