package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/CrossoversForCures/Tournament-Scoring/models"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Home page</h1>")
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Teams!!</h1>")
}

// func main() {
// 	http.HandleFunc("/", homeHandler)
// 	http.HandleFunc("/teams", teamsHandler)

// 	fmt.Println("Starting the server on port 3000...")
// 	http.ListenAndServe(":3000", nil)
// }

func main() {
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

	coll := client.Database("tournament_scoring").Collection("tournaments")
	newTournament := models.Tournament{Name: "E4E 2024", Date: time.Now(), Status: "Active"}

	result, err := coll.InsertOne(context.TODO(), newTournament)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Document inserted with ID: %s\n", result.InsertedID)
}
