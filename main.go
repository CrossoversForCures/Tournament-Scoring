package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/CrossoversForCures/Tournament-Scoring/models"
	"github.com/joho/godotenv"
)

var coll *mongo.Collection
var err error

func tournamentHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{Key: "status", Value: "Active"}}
	var result models.Tournament
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			errorResponse := map[string]string{"error": "No available tournaments found"}
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		panic(err)
	}

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	w.Write(output)
}

func tournamentStartHandler(w http.ResponseWriter, r *http.Request) {
	filter := bson.D{{Key: "status", Value: "Registering"}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: "Active"}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		panic(err)
	}
	if result.ModifiedCount != 1 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No available tournaments to start"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
}
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	uri := os.Getenv("DATABASE_URL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll = client.Database("tournament_scoring").Collection("tournaments")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/tournament", tournamentHandler)
	mux.HandleFunc("POST /api/tournament/start", tournamentStartHandler)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", mux)
	// 	newTournament := models.Tournament{Name: "E4E 2024", Date: time.Now(), Status: "Active"}

	// 	result, err := coll.InsertOne(context.TODO(), newTournament)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}
