package main

import ( 
	"net/http"
	"sync"
    "encoding/json"
	"strings"
)


var (
    favorites = make(map[string]bool) // (dict in python) to store favorite Pokémon. The key is the Pokémon name, and the value is a boolean (true if it's a favorite).
	mu sync.Mutex //A mutex to ensure that only one goroutine can access the favorites map at a time, preventing race conditions.

)


// GET
func getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/pokemon/")
	if name ==""{
		var allPokemons []Pokemon
		for _, value := range POKEDEX {
			// _ is the key, value the value of the """dict"""
			allPokemons = append(allPokemons, value)
		}
		json.NewEncoder(w).Encode(allPokemons)
		return
	}

	pokemon, exists := POKEDEX[name]
	if !exists {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pokemon)
}

// Post
func postFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	var pokemon Pokemon
	// Decodes the JSON body of the request into the pokemon variable.
	if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	favorites[pokemon.Name] = true
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
}

// Put
func putDeleteFavoriteHandler(w http.ResponseWriter, r *http.Request) {
	var pokemon Pokemon
	// Decode the JSON body of the request into the pokemon variable.
	if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}
	mu.Lock()
	favorites[pokemon.Name] = False
	mu.Unlock()
	w.WriteHeader(http)
}