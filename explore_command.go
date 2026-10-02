package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		fmt.Println("Please provide a name or id of a location you want to explore")
		return nil
	}
	jsonBody, err := cfg.client.LocationExploreRequest(args[0])
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\nFound Pokemon:\n", jsonBody.Name)
	for _, encounter := range jsonBody.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
