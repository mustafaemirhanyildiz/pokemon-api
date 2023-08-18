package routes

import (
	"github.com/gorilla/mux"
	"pokemon-api/controllers"
)

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/api/pokemons", controllers.PokemonListEndpoint).Methods("GET")
	// ... diğer rotaları burada tanımlayabilirsiniz
}
