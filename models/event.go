package models

import (
	"context"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name   string             `bson:"name,omitempty" json:"name,omitempty"`
	Status int                `bson:"status,omitempty" json:"status,omitempty"`
	//0: In registration | 1: Playing pools | 2: Playing bracket | 3: Completed
	PoolRounds int `bson:"poolRounds,omitempty" json:"poolRounds,omitempty"`
}

func GetEvent(_id primitive.ObjectID) Event {
	var result Event
	err := configs.EventsCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func GetEvents() []Event {
	cursor, err := configs.EventsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []Event
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func UpdateEvent(_id primitive.ObjectID, update bson.D) {
	_, err := configs.EventsCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}, update)
	if err != nil {
		panic(err)
	}
}

// Test method
func AddEvents() {
	_, err := configs.EventsCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

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

	_, err = configs.EventsCollection.InsertMany(context.TODO(), newEvents)
	if err != nil {
		panic(err)
	}
}
