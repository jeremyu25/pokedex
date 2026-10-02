package main

import (
	"time"

	"github.com/jeremyu25/pokedex/internal/pokeapi"
)

func main() {
	config := &config{
		commands: getCommands(),
		client:   pokeapi.NewClient(5*time.Second, 5*time.Second),
		pokedex:  map[string]pokeapi.PokemonJson{},
	}
	startRepl(config)
}
