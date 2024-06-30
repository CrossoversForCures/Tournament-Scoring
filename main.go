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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Hi</h1>")
	filter := bson.D{{Key: "status", Value: "Active"}}
	var result models.Tournament
	err = coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {

			return
		}
		panic(err)
	}

	output, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", output)
	fmt.Fprintf(w, "<h1>Welcome to %v!</h1>", result.Name)
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Teams</h1>")
}

var coll *mongo.Collection
var err error

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Making database connection
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

	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", nil)
	// 	newTournament := models.Tournament{Name: "E4E 2024", Date: time.Now(), Status: "Active"}

	// 	result, err := coll.InsertOne(context.TODO(), newTournament)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}
