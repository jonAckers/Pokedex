package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (PokemonResp, error) {
	url := baseUrl + "pokemon/" + pokemon + "/"

	cached, exists := c.cache.Get(url)
	if exists {
		pokemonResp := PokemonResp{}
		err := json.Unmarshal(cached, &pokemonResp)
		if err != nil {
			return PokemonResp{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	date, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	c.cache.Add(url, date)

	pokemonResp := PokemonResp{}
	err = json.Unmarshal(date, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	return pokemonResp, nil
}
