# Pokemon API

A simple Go API to retrieve and manage Pokémon data. The API allows users to get details of specific Pokémon, add Pokémon to their favorites, and remove them from favorites.

## Features

- Retrieve details of a specific Pokémon by name.
- Retrieve details of all Pokémon.
- Add a Pokémon to favorites.
- Remove a Pokémon from favorites.

## Getting Started

### Prerequisites

- Go (version 1.16+ recommended)

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/pokemon-api.git
   cd pokemon-api
   ```

2. Initialize the Go module:

    ```sh

    go mod init pokemon-api
    ```

3. Run the server:

    ```sh

        go run main.go handlers.go pokemon.go
    ``` 
The server will start on http://localhost:8080.

## API Endpoints
1. Retrieve All Pokémon
    ```
    URL: /pokemon/
    Method: GET
    Response:
    [
        {"name": "Pikachu", "type": "Electric"},
        {"name": "Bulbasaur", "type": "Grass/Poison"}
        // Add more Pokémon as needed
    ]
    ``` 

2. Retrieve a Specific Pokémon
    ```
    URL: /pokemon/{name}
    Method: GET
    Response:
    {
        "name": "Pikachu",
        "type": "Electric"
    }

    or

    Pokemon not found
    ```
3. Add a Pokémon to Favorites

    URL: /favorites
    Method: POST
    Request Body:

    {
        "name": "Pikachu"
    }

    Response: 201 Created

## Project Structure
``` plaintext
.
├── main.go
├── handlers.go
└── pokemon.go

    main.go: Initializes the server, defines global variables, and sets up routes.
    handlers.go: Contains the handler functions for the API endpoints.
    pokemon.go: Defines the Pokemon struct and the simulated database of Pokémon.
```