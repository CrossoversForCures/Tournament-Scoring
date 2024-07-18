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
var BracketsCollection *mongo.Collection

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
	PoolGamesCollection = client.Database("tournament_scoring").Collection("pools")
	BracketsCollection = client.Database("tournament_scoring").Collection("brackets")
}

func GetStripeKey() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	key := os.Getenv("STRIPE_SECRET_KEY")
	if key == "" {
		log.Fatal("STRIPE_API_KEY is not set in the environment")
	}
	return key
}
