package models

import (
	"context"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Event       primitive.ObjectID `bson:"event_id"`
	PoolsPlayed int                `bson:"pools_played"`
}

func AddTeam() {
	newTeam := Team{Name: "Team1"}
	_, err := configs.GetCollection(configs.DB, "teams").InsertOne(context.TODO(), newTeam)
	if err != nil {
		panic(err)
	}
}
