package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/pokemon/", getPokemonHandler)
    http.HandleFunc("/favorites", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPut {
            //putFavoriteHandler(w, r)
			fmt.Println("Not implemented lol")
        } else if r.Method == http.MethodPost {
            postFavoriteHandler(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}