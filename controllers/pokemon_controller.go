package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	// "github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"../models"
)

var pokemonCollection *mongo.Collection

func init() {
	pokemonCollection = client.Database("pokemonDB").Collection("pokemons")
}

// PokemonListEndpoint tüm Pokemonları listeler
func PokemonListEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")

	var pokemons []models.Pokemon
	cur, err := pokemonCollection.Find(context.Background(), bson.D{}, options.Find())
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message": "Veritabanından Pokemonlar getirilirken bir hata oluştu."}`))
		return
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var pokemon models.Pokemon
		cur.Decode(&pokemon)
		pokemons = append(pokemons, pokemon)
	}

	json.NewEncoder(response).Encode(pokemons)
}

// ... diğer controller fonksiyonlarını burada tanımlayabilirsiniz
