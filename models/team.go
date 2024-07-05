package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string
	TournamentID primitive.ObjectID `bson:"tournament_id"`
	EventID      primitive.ObjectID `bson:"event_id"`
	Player1      string
	Player2      string
	Player3      string
	Group        string
	Games        []primitive.ObjectID
	Seed         int
	PostSeed     int `bson:"post_seed"`
	Final        int
}
