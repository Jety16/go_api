package main

import (
	"fmt"  // Provides I/O functions like Println.
	"net/http"  // Provides HTTP client and server implementations.
	"encoding/json"  // Used to encode and decode JSON.
	"strings"  // Contains string manipulation functions.
	"sync"  // Provides basic synchronization primitives like mutexes
)

var (
    favorites = make(map[string]bool) // (dict in python) to store favorite Pokémon. The key is the Pokémon name, and the value is a boolean (true if it's a favorite).
	mu sync.Mutex //A mutex to ensure that only one goroutine can access the favorites map at a time, preventing race conditions.

)


func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, WOrLd!")
}

func main() {
	// Route setup
	http.HandleFunc("/", handler)
	fmt.Println("Starting Server at port 8080...")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}