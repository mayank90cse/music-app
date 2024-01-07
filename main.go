package main

import (
	"context"
	"fmt"
	"log"
	"music-app/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := mux.NewRouter()
	dbClient := GetMongoClient()
	ctx := handlers.New(dbClient)

	router.HandleFunc("/api/v1/metadata", ctx.CreateMusicData).Methods("POST")
	router.HandleFunc("/api/v1/track/{isrc}", ctx.FetchMusicByIsrc).Methods("GET")
	router.HandleFunc("/api/v1/artist/track", ctx.FetchMusicByArtist).Methods("GET")

	startServer(router)
}

func startServer(router *mux.Router) {
	port := os.Getenv("APP_SERVER_PORT")
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println("Error while starting server", err)
		return
	}
	log.Printf("Server started and listening on port :" + port)
}

func GetMongoClient() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	username := os.Getenv("APP_MONGO_DB_USERNAME")
	pwd := os.Getenv("APP_MONGO_DB_PASSWORD")
	opts := options.Client().ApplyURI("mongodb+srv://" + username + ":" + pwd + "@cluster.xhlj8qj.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	/*	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}() */
	// Send a ping to confirm a successful connection
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}
