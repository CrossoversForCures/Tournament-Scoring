package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty" validate:"required"`
	Name  string               `bson:"name,omitempty" validate:"required"`
	Teams []primitive.ObjectID `bson:"teams,omitempty" validate:"required"`
	Games []primitive.ObjectID `bson:"games,omitempty" validate:"required"`
}

func AddEvents() {
	newEvents := []interface{}{
		Event{Name: "3rd/4th Boys"},
		Event{Name: "5th/6th Boys"},
		Event{Name: "7th/8th Boys"},
		Event{Name: "9th/10th Boys"},
		Event{Name: "11th/12th Boys"},
		Event{Name: "5th/6th Girls"},
		Event{Name: "7th/8th Girls"},
		Event{Name: "College Co-ed"},
	}

	_, err := EventsCollection.InsertMany(context.TODO(), newEvents)
	if err != nil {
		panic(err)
	}
}
