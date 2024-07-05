package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string
	Teams []primitive.ObjectID `bson:"teams"`
	Games []primitive.ObjectID `bson:"games"`
}
