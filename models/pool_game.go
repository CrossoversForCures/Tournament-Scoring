package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolGame struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Event primitive.ObjectID `bson:"event_id"`
	Round int                `bson:"round, omitempty"`
	Team1 Team               `bson:"team1, omitempty"`
	Team2 Team               `bson:"team2, omitempty"`
}
