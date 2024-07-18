package models

import (
	"context"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Event       string             `bson:"event,omitempty" json:"event,omitempty"`
	PoolsWon    int                `bson:"poolsWon,omitempty" json:"poolsWon,omitempty"`
	TotalPoints int                `bson:"totalPoints,omitempty" json:"totalPoints,omitempty"`
	Seeding     int                `bson:"seeding,omitempty" json:"seeding,omitempty"`

	poolsPlayed int
}

func GetTeam(_id primitive.ObjectID) Team {
	var result Team
	err := configs.TeamsCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func GetTeams(event string) []Team {
	cursor, err := configs.TeamsCollection.Find(context.TODO(), bson.D{{Key: "event", Value: event}})
	if err != nil {
		panic(err)
	}

	var results []Team
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func UpdateTeam(_id primitive.ObjectID, update bson.D) {
	_, err := configs.TeamsCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}, update)
	if err != nil {
		panic(err)
	}
}

// Test Method
func AddTeams() {
	_, err := configs.TeamsCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	_, err = configs.PoolGamesCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	newTeams := []interface{}{
		Team{Name: "Team A", Event: "3rd-4th-boys"},
		Team{Name: "Team B", Event: "3rd-4th-boys"},
		Team{Name: "Team C", Event: "3rd-4th-boys"},
		Team{Name: "Team D", Event: "3rd-4th-boys"},
		Team{Name: "Team E", Event: "3rd-4th-boys"},
		Team{Name: "Team F", Event: "3rd-4th-boys"},
		Team{Name: "Team G", Event: "3rd-4th-boys"},
		Team{Name: "Team H", Event: "3rd-4th-boys"},
		Team{Name: "Team I", Event: "3rd-4th-boys"},
		Team{Name: "Team J", Event: "3rd-4th-boys"},
	}

	_, err = configs.TeamsCollection.InsertMany(context.TODO(), newTeams)

	if err != nil {
		panic(err)
	}
}
