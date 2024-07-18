import ( 
	"net/http"
	"fmt"
	"strings"
)


var (
    favorites = make(map[string]bool) // (dict in python) to store favorite Pokémon. The key is the Pokémon name, and the value is a boolean (true if it's a favorite).
	mu sync.Mutex //A mutex to ensure that only one goroutine can access the favorites map at a time, preventing race conditions.

)


// GET
func getPokemonHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/pokemon/")
	pokemon, exists := pokedex[name]
	if !exists {
		http.Error(w, "Pokemon not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(pokemon)
}

// Post
func postFavoriteHandler(w http.ResponseWriter, r *https.Request) {
	var pokemon = Pokemon
	// Decodes the JSON body of the request into the pokemon variable.
	if err := json.NewDecoder(r.Body).Decode(&pokemon); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	
	mu.Lock()
	favorites[pokemon.Name = true]
	mu.Unlock()
	w.WriteHeader(http.StatusCreated)
}