package models

import (
	"context"
	"time"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name       string             `bson:"name,omitempty" json:"name,omitempty"`
	Slug       string             `bson:"slug,omitempty" json:"slug,omitempty"`
	Time       time.Time          `bson:"time,omitempty" json:"time,omitempty"`
	Status     int                `bson:"status,omitempty" json:"status,omitempty"` //0: In registration | 1: Playing pools | 2: Playing bracket | 3: Completed
	PoolRounds int                `bson:"poolRounds,omitempty" json:"poolRounds,omitempty"`
}

func GetEvent(event string) Event {
	var result Event
	err := configs.EventsCollection.FindOne(context.TODO(), bson.D{{Key: "slug", Value: event}}).Decode(&result)
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

func UpdateEvent(eventSlug string, update bson.D) {
	_, err := configs.EventsCollection.UpdateOne(context.TODO(), bson.D{{Key: "slug", Value: eventSlug}}, update)
	if err != nil {
		panic(err)
	}
}

func InitiateEvents() {
	_, err := configs.EventsCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	newEvents := []interface{}{
		Event{Name: "3rd/4th Boys", Slug: "3rd-4th-boys"},
		Event{Name: "5th/6th Boys", Slug: "5th-6th-boys"},
		Event{Name: "7th/8th Boys", Slug: "7th-8th-boys"},
		Event{Name: "9th/10th Boys", Slug: "9th-10th-boys"},
		Event{Name: "11th/12th Boys", Slug: "11th-12th-boys"},
		Event{Name: "3rd/4th Girls", Slug: "3rd-4th-girls"},
		Event{Name: "5th/6th Girls", Slug: "5th-6th-girls"},
		Event{Name: "7th/8th Girls", Slug: "7th-8th-girls"},
		Event{Name: "9th/12th Girls", Slug: "9th-12th-girls"},
		Event{Name: "College Co-ed", Slug: "college-coed"},
	}

	_, err = configs.EventsCollection.InsertMany(context.TODO(), newEvents)
	if err != nil {
		panic(err)
	}
}
