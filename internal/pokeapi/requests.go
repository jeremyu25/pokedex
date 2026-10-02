package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	locationAreaUrl = "https://pokeapi.co/api/v2/location-area/"
	pokemonUrl      = "https://pokeapi.co/api/v2/pokemon/"
)

func (c *Client) LocationAreaRequest(requestUrl *string) (LocationAreaJson, error) {
	jsonBody := LocationAreaJson{}
	url := locationAreaUrl
	if requestUrl != nil {
		url = *requestUrl
	}
	cachedData, found := c.cache.Get(url)
	if found {
		if err := json.Unmarshal(cachedData, &jsonBody); err != nil {
			return LocationAreaJson{}, err
		}
		return jsonBody, nil

	} else {
		req, err := http.NewRequest("GET", url, nil)
		res, err := c.httpClient.Do(req)

		if err != nil {
			return LocationAreaJson{}, err
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaJson{}, err
		}
		defer res.Body.Close()
		if err := json.Unmarshal(data, &jsonBody); err != nil {
			return LocationAreaJson{}, err
		}
		c.cache.Add(url, data)
		return jsonBody, nil
	}

}

func (c *Client) LocationExploreRequest(location string) (LocationExploreJson, error) {
	jsonBody := LocationExploreJson{}
	url := locationAreaUrl + location
	cachedData, found := c.cache.Get(url)
	if found {
		if err := json.Unmarshal(cachedData, &jsonBody); err != nil {
			return LocationExploreJson{}, err
		}
		return jsonBody, nil

	} else {
		req, err := http.NewRequest("GET", url, nil)
		res, err := c.httpClient.Do(req)

		if err != nil {
			return LocationExploreJson{}, err
		}
		if res.StatusCode == 404 {
			return LocationExploreJson{}, fmt.Errorf("Location does not exist")
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return LocationExploreJson{}, err
		}
		defer res.Body.Close()
		if err := json.Unmarshal(data, &jsonBody); err != nil {
			return LocationExploreJson{}, err
		}
		c.cache.Add(url, data)
		return jsonBody, nil
	}
}
func (c *Client) GetPokemonRequest(pokemon string) (PokemonJson, error) {
	jsonBody := PokemonJson{}
	url := pokemonUrl + pokemon
	cachedData, found := c.cache.Get(url)
	if found {
		if err := json.Unmarshal(cachedData, &jsonBody); err != nil {
			return PokemonJson{}, err
		}
		return jsonBody, nil

	} else {
		req, err := http.NewRequest("GET", url, nil)
		res, err := c.httpClient.Do(req)

		if err != nil {
			return PokemonJson{}, err
		}
		if res.StatusCode == 404 {
			return PokemonJson{}, fmt.Errorf("Pokemon does not exist")
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return PokemonJson{}, err
		}
		defer res.Body.Close()
		if err := json.Unmarshal(data, &jsonBody); err != nil {
			return PokemonJson{}, err
		}
		c.cache.Add(url, data)
		return jsonBody, nil
	}
}
