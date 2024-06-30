package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	newTournament := Tournament{}
	title := "Back to the Future"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).
		Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
