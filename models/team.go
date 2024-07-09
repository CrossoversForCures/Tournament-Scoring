package models

import (
	"context"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Event       primitive.ObjectID `bson:"event_id"`
	PoolsPlayed int                `bson:"pools_played"`
}

// Test Method
func AddTeams() {
	_, err := configs.TeamsCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	eventId, _ := primitive.ObjectIDFromHex("668c8028b526a0689257b27e")

	newTeams := []interface{}{
		Team{Name: "Team A", Event: eventId},
		Team{Name: "Team B", Event: eventId},
		Team{Name: "Team C", Event: eventId},
		Team{Name: "Team D", Event: eventId},
	}

	_, err = configs.TeamsCollection.InsertMany(context.TODO(), newTeams)
	if err != nil {
		panic(err)
	}
}
