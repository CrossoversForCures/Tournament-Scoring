package models

import (
	"context"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Event       primitive.ObjectID `bson:"eventId,omitempty" json:"eventId,omitempty"`
	poolsPlayed int
	PoolsWon    int `bson:"poolsWon,omitempty" json:"poolsWon,omitempty"`
	TotalPoints int `bson:"totalPoints,omitempty" json:"totalPoints,omitempty"`
	Seeding     int `bson:"seeding,omitempty" json:"seeding,omitempty"`
}

func GetTeam(_id primitive.ObjectID) Team {
	var result Team
	err := configs.TeamsCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func GetTeams(eventId primitive.ObjectID) []Team {
	cursor, err := configs.TeamsCollection.Find(context.TODO(), bson.D{{Key: "eventId", Value: eventId}})
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

	var result Event
	err = configs.EventsCollection.FindOne(context.TODO(), bson.D{{Key: "name", Value: "3rd/4th Boys"}}).Decode(&result)
	if err != nil {
		panic(err)
	}
	eventId, _ := primitive.ObjectIDFromHex(result.ID.Hex())

	newTeams := []interface{}{
		Team{Name: "Team A", Event: eventId},
		Team{Name: "Team B", Event: eventId},
		Team{Name: "Team C", Event: eventId},
		Team{Name: "Team D", Event: eventId},
		Team{Name: "Team E", Event: eventId},
		Team{Name: "Team F", Event: eventId},
		Team{Name: "Team G", Event: eventId},
		Team{Name: "Team H", Event: eventId},
	}

	_, err = configs.TeamsCollection.InsertMany(context.TODO(), newTeams)

	if err != nil {
		panic(err)
	}
}
