package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string
	EventID  primitive.ObjectID `bson:"event_id"`
	Player1  string
	Player2  string
	Player3  string
	Group    string
	Games    []primitive.ObjectID
	Seed     int
	PostSeed int `bson:"post_seed"`
	Final    int
}

func AddTeam() {
	newTeam := Team{Name: "Team1"}
	// Inserts sample documents into the collection
	_, err := EventsCollection.InsertOne(context.TODO(), newTeam)
	if err != nil {
		panic(err)
	}
}
