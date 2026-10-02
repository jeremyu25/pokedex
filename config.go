package main

import (
	"github.com/jeremyu25/pokedex/internal/pokeapi"
	"github.com/jeremyu25/pokedex/internal/pokecache"
)

type config struct {
	commands map[string]cliCommand
	next     *string
	prev     *string
	client   pokeapi.Client
	cache    pokecache.Cache
	pokedex  map[string]pokeapi.PokemonJson
}
