package models

import (
	"context"
	"math/rand"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolGame struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Event      string             `bson:"event,omitempty" json:"event,omitempty"`
	Round      int                `bson:"round,omitempty" json:"round,omitempty"`
	Court      string             `bson:"court,omitempty" json:"court,omitempty"`
	Status     int                `bson:"status,omitempty" json:"status,omitempty"`
	Team1      primitive.ObjectID `bson:"team1Id,omitempty" json:"team1Id,omitempty"`
	Team2      primitive.ObjectID `bson:"team2Id,omitempty" json:"team2Id,omitempty"`
	Team1Name  string             `bson:"team1Name,omitempty" json:"team1Name,omitempty"`
	Team2Name  string             `bson:"team2Name,omitempty" json:"team2Name,omitempty"`
	Team1Score int                `bson:"team1Score,omitempty" json:"team1Score,omitempty"`
	Team2Score int                `bson:"team2Score,omitempty" json:"team2Score,omitempty"`
}

func GetPool(_id primitive.ObjectID) PoolGame {
	var result PoolGame
	err := configs.PoolGamesCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func GetPools(event string) []PoolGame {
	cursor, err := configs.PoolGamesCollection.Find(context.TODO(), bson.D{{Key: "event", Value: event}})
	if err != nil {
		panic(err)
	}

	var results []PoolGame
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}

	return results
}

func UpdatePool(_id primitive.ObjectID, update bson.D) {
	_, err := configs.PoolGamesCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}, update)
	if err != nil {
		panic(err)
	}
}

func InsertPool(newPool PoolGame) {
	_, err := configs.PoolGamesCollection.InsertOne(context.TODO(), newPool)
	if err != nil {
		panic(err)
	}
}

func SortPools(event string) {
	_, err := configs.PoolGamesCollection.DeleteMany(context.TODO(), bson.D{{Key: "event", Value: event}})
	if err != nil {
		panic(err)
	}
	teams := GetTeams(event)

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

	gamesPlayed := make([]matchup, 0)

	// Keep doing rounds until all teams have played 2 games
	for len(teams) != 0 {
		teamsLeftThisRound := min(teamsPerRound, len(teams))

		var currentGames []matchup

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
				currentGames = append(currentGames, matchup{Team1: team1, Team2: team2, Round: round})
			}
		}

		// Mark this round's games as played
		gamesPlayed = append(gamesPlayed, currentGames...)
		// Remove this rounds teams from the queue
		teams = teams[teamsLeftThisRound:]
		for _, game := range currentGames {
			game.Team1.poolsPlayed++
			game.Team2.poolsPlayed++

			// If the team isn't finished with pools, add it back in
			if game.Team1.poolsPlayed < 2 {
				teams = append(teams, game.Team1)
			}
			if game.Team2.poolsPlayed < 2 {
				teams = append(teams, game.Team2)
			}
		}
		round++
	}

	// Add all the games to the database
	for i, game := range gamesPlayed {
		courts := [4]string{"A", "B", "C", "D"}
		newGame := PoolGame{
			Event:     event,
			Round:     game.Round,
			Court:     courts[i%len(courts)],
			Team1:     game.Team1.ID,
			Team2:     game.Team2.ID,
			Team1Name: game.Team1.Name,
			Team2Name: game.Team2.Name,
		}
		InsertPool(newGame)
	}
}

func duplicateGame(games []matchup, target []matchup) bool {
	for _, game := range target {
		for _, item := range games {
			if item.Team1.Name == game.Team1.Name && item.Team2.Name == game.Team2.Name {
				return true
			}
		}
	}

	return false
}

type matchup struct {
	Team1 Team
	Team2 Team
	Round int
}
