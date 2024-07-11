package configs

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var EventsCollection *mongo.Collection
var TeamsCollection *mongo.Collection
var PoolGamesCollection *mongo.Collection

func ConnectDB() {
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

	EventsCollection = client.Database("tournament_scoring").Collection("events")
	TeamsCollection = client.Database("tournament_scoring").Collection("teams")
	PoolGamesCollection = client.Database("tournament_scoring").Collection("games")
}
