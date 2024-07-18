package main

type Pokemon struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

// A map simulating a database of Pokémon.
// The key is the Pokémon name (in lowercase), and the value is a Pokemon struct.
var POKEDEX = map[string]Pokemon{
    "pikachu": {"Pikachu", "Electric"},
    "bulbasaur": {"Bulbasaur", "Grass/Poison"},
}