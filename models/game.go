package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Team1 Team               `bson:"team1, omitempty"`
	Team2 Team               `bson:"team2, omitempty"`
}
