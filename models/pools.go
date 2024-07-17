package models

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PoolGame struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Event      primitive.ObjectID `bson:"eventId,omitempty" json:"eventId,omitempty"`
	Round      int                `bson:"round,omitempty" json:"round,omitempty"`
	Team1      primitive.ObjectID `bson:"team1Id,omitempty" json:"team1Id,omitempty"`
	Team2      primitive.ObjectID `bson:"team2Id,omitempty" json:"team2Id,omitempty"`
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

func GetPools(eventId primitive.ObjectID) []PoolGame {
	cursor, err := configs.PoolGamesCollection.Find(context.TODO(), bson.D{{Key: "eventId", Value: eventId}})
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

type matchup struct {
	Team1 Team
	Team2 Team
	Round int
}

func SortPools(eventId primitive.ObjectID) {
	_, err := configs.PoolGamesCollection.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	teams := GetTeams(eventId)

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
		fmt.Printf("Round: %v\n", round)

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
			fmt.Printf("%v vs %v\n", game.Team1.Name, game.Team2.Name)

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
	for _, game := range gamesPlayed {
		newGame := PoolGame{
			Event: eventId,
			Round: game.Round,
			Team1: game.Team1.ID,
			Team2: game.Team2.ID,
		}
		InsertPool(newGame)
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "poolRounds", Value: round - 1}}}}
	UpdateEvent(eventId, update)
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