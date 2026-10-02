# Pokedex Project

Simple REPL project that utilises the [PokeApi](https://pokeapi.co/) to explore locations in the Pokemon world with the simple functionality of catching Pokemon and adding them to a Pokedex. This project implements simple time-based caching of api results using a predefined interval for smoother loading of reused commands

## Commands

### inspect \<name of pokemon>
Inspect a caught Pokemon
### pokedex
List the name of all Pokemon in user's pokedex
### exit
Exit the Pokedex

### help
Displays a help message

### map
Displays 20 location areas in Pokemon World. Reusing this command will load the next 20 location areas.

### mapb
Displays 20 previous location areas in Pokemon World

### explore \<name of location area>
Displays list of Pokemon in location area

### catch \<name of pokemon>
Attempts to catch a specified Pokemon

# Acknowledgements
## Pokemon API: https://pokeapi.co/
