package models

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID    primitive.ObjectID   `bson:"_id,omitempty" validate:"required"`
	Name  string               `bson:"name,omitempty" validate:"required"`
	Teams []primitive.ObjectID `bson:"teams,omitempty" validate:"required"`
	Games []primitive.ObjectID `bson:"games,omitempty" validate:"required"`
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

func SortPools(eventId primitive.ObjectID) {
	// Get list of teams from dadtabase
	cursor, err := configs.TeamsCollection.Find(context.TODO(), bson.M{"event_id": eventId})
	var teams []Team
	if err = cursor.All(context.TODO(), &teams); err != nil {
		panic(err)
	}

	numTeams := len(teams)
	round := 1

	var teamsPerRound int
	if numTeams >= 8 {
		teamsPerRound = 8
	} else if numTeams >= 6 {
		teamsPerRound = 6
	} else if numTeams >= 4 {
		teamsPerRound = 4
	} else {
		teamsPerRound = 2
	}

	gamesPlayed := make([]PoolGame, 0)

	// Keep doing rounds until all teams have played 2 games
	for len(teams) != 0 {
		teamsLeftThisRound := min(teamsPerRound, len(teams))
		fmt.Printf("Round: %v\n", round)

		var currentGames []PoolGame

		for ok := true; ok; ok = (duplicateGame(gamesPlayed, currentGames)) {
			currentGames = nil
			// Shuffle the teams left this round
			for i := 0; i < teamsLeftThisRound; i++ {
				j := rand.Intn(i + 1)
				teams[i], teams[j] = teams[j], teams[i]
			}

			for i := 0; i < teamsLeftThisRound; i += 2 {
				team1 := teams[i]
				team2 := teams[i+1]
				// Make sure Team 1 is alphabetically lower
				if team2.Name < team1.Name {
					team1, team2 = team2, team1
				}
				currentGames = append(currentGames, PoolGame{Team1: team1, Team2: team2, Round: round, Event: eventId})
			}
		}
		// Mark this round's games as played
		gamesPlayed = append(gamesPlayed, currentGames...)
		// Remove this rounds teams from the queue
		teams = teams[teamsLeftThisRound:]
		for _, game := range currentGames {
			game.Team1.PoolsPlayed++
			game.Team2.PoolsPlayed++
			fmt.Printf("%v vs %v\n", game.Team1.Name, game.Team2.Name)

			// If the team isn't finished with pools, add it back in
			if game.Team1.PoolsPlayed < 2 {
				teams = append(teams, game.Team1)
			}
			if game.Team2.PoolsPlayed < 2 {
				teams = append(teams, game.Team2)
			}
		}
		round++
	}

	// Add all the games to the database
	for _, game := range gamesPlayed {

		_, err := configs.GamesCollection.InsertOne(context.TODO(), game)
		if err != nil {
			panic(err)
		}
	}
}

func duplicateGame(games []PoolGame, target []PoolGame) bool {
	for _, game := range target {
		for _, item := range games {
			if item.Team1.Name == game.Team1.Name && item.Team2.Name == game.Team2.Name {
				return true
			}
		}
	}

	return false
}
