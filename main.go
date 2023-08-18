package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger" // Swagger UI için import
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	// MongoDB bağlantısını kur
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.Background(), clientOptions)

	// HTTP sunucusunu başlat
	router := mux.NewRouter()

	// Swagger UI'ı etkinleştir (Bu kısmı HTTP sunucusunu başlatmadan önce eklemelisiniz)
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // Swagger JSON dosyasının URL'si
	))

	log.Fatal(http.ListenAndServe(":8080", router))
}
