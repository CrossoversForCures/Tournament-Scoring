package models

import (
	"context"
	"fmt"
	"log"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

type Team struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name        string             `bson:"name,omitempty" json:"name,omitempty"`
	Event       string             `bson:"event,omitempty" json:"event,omitempty"`
	PoolsWon    int                `bson:"poolsWon,omitempty" json:"poolsWon,omitempty"`
	TotalPoints int                `bson:"totalPoints,omitempty" json:"totalPoints,omitempty"`
	Seeding     int                `bson:"seeding,omitempty" json:"seeding,omitempty"`
	Rank        int                `bson:"rank,omitempty" json:"rank,omitempty"`

	poolsPlayed int
}

func GetTeam(_id primitive.ObjectID) Team {
	var result Team
	err := configs.TeamsCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func GetTeams(event string) []Team {
	cursor, err := configs.TeamsCollection.Find(context.TODO(), bson.D{{Key: "event", Value: event}})
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

func InsertTeam(newTeam Team) {
	_, err := configs.TeamsCollection.InsertOne(context.TODO(), newTeam)
	if err != nil {
		panic(err)
	}
}

// Add all teams
func AddTeams() {
	_, err := configs.TeamsCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	_, err = configs.PoolGamesCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	newTeams, err := GetAllTeams()
	if err != nil {
		panic(err)
	}
	newTeamsInterface := make([]interface{}, len(newTeams))
	for i, t := range newTeams {
		newTeamsInterface[i] = t
	}
	_, err = configs.TeamsCollection.InsertMany(context.TODO(), newTeamsInterface)

	if err != nil {
		panic(err)
	}
}

func GetAllTeams() ([]Team, error) {
	// Initialize params for listing checkout sessions
	params := &stripe.CheckoutSessionListParams{}

	// Create a slice to store all sessions
	var teamRegistrations []Team

	// Use an iterator to paginate through all checkout sessions
	i := session.List(params)
	for i.Next() {
		s := i.CheckoutSession()
		if s.CustomFields != nil && len(s.CustomFields) > 2 {
			team, err := processCheckoutSessionTest(s)
			if err != nil {
				continue
			}
			if team.Event != "" {
				teamRegistrations = append(teamRegistrations, team)
			}
		}

	}
	// Handle any error encountered during the iteration
	if err := i.Err(); err != nil {
		log.Fatalf("Error listing sessions: %v", err)
	}

	fmt.Printf("Total sessions retrieved: %d\n", len(teamRegistrations))
	return teamRegistrations, nil
}

func processCheckoutSessionTest(s *stripe.CheckoutSession) (Team, error) {
	team := Team{}

	for _, field := range s.CustomFields {
		switch field.Key {
		case "teamname":
			if field.Text != nil {
				team.Name = field.Text.Value
			}
		case "division":
			if field.Dropdown != nil {
				event := field.Dropdown.Value
				if event == "5th6thgirls" {
					team.Event = "5th-6th-girls"
				}
				if event == "5th6thboys" {
					team.Event = "5th-6th-boys"
				}
				if event == "7th8thboys" {
					team.Event = "7th-8th-boys"
				}
			}
		}
	}
	return team, nil
}

func InitiateTestTeams() {
	_, err := configs.TeamsCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	_, err = configs.PoolGamesCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	newTeams := []interface{}{
		Team{
			Name:  "Team A",
			Event: "5th-6th-boys",
		},
		Team{
			Name:  "Team B",
			Event: "5th-6th-boys",
		},
		Team{
			Name:  "Team C",
			Event: "5th-6th-boys",
		},
		Team{
			Name:  "Team D",
			Event: "5th-6th-boys",
		},
		Team{
			Name:  "Team E",
			Event: "5th-6th-boys",
		},
	}

	_, err = configs.TeamsCollection.InsertMany(context.TODO(), newTeams)
	if err != nil {
		log.Fatal(err)
	}
}
