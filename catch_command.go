package main

import (
	"fmt"
	"math/rand/v2"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide the name of a Pokemon you want to catch")
		return nil
	}
	pokemon := args[0]
	jsonBody, err := cfg.client.GetPokemonRequest(pokemon)
	if err != nil {
		return err
	}
	pokemon = jsonBody.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	catchRate := 70 - (jsonBody.BaseExperience / 8)
	if rand.IntN(100) >= catchRate {
		fmt.Printf("%s escaped!\n", pokemon)
		return nil
	} else {
		cfg.pokedex[pokemon] = jsonBody
		fmt.Printf("%s was caught!\n", pokemon)
		return nil
	}
}
