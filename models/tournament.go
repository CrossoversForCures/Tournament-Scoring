package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tournament model
type Tournament struct {
	ID     primitive.ObjectID   `bson:"_id,omitempty"`
	Name   string               `bson:"name"`
	Date   time.Time            `bson:"date"`
	Status string               `bson:"status"`
	Teams  []primitive.ObjectID `bson:"teams"`
	Games  []primitive.ObjectID `bson:"games"`
}
