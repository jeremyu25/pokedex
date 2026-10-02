package main

import (
	"fmt"
	"strings"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("Please provide a single Pokemon name to inspect")
		return nil
	}
	pokemonName := strings.ToLower(args[0])
	pokemon, found := cfg.pokedex[pokemonName]
	if !found {
		fmt.Println("You have not caught that Pokemon")
	} else {
		fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
`, pokemon.Name, pokemon.Height, pokemon.Weight)

		for _, stat := range pokemon.Stats {
			fmt.Printf("-%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("-%s\n", pokemonType.Type.Name)
		}
	}
	return nil
}
