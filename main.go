package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"github.com/CrossoversForCures/Tournament-Scoring/models"
)

func startHandler(w http.ResponseWriter, r *http.Request) {

}

func sortPools() {
	duplicateGame := func(games []models.Game, target []models.Game) bool {
		for _, game := range target {
			for _, item := range games {
				if item.Team1.Name == game.Team1.Name && item.Team2.Name == game.Team2.Name {
					return true
				}
			}
		}

		return false
	}

	teams := make([]models.Team, 0)
	numTeams := 5
	round := 1

	teamsPerRound := 4

	gamesPlayed := make([]models.Game, 0)

	//Creating the teams for testing
	for i := 0; i < numTeams; i++ {
		newName := 'A' + i
		teams = append(teams, models.Team{Name: string(newName), PoolsPlayed: 0})
	}

	// Keep doing rounds until all teams have played 2 games
	for len(teams) != 0 {
		teamsLeftThisRound := min(teamsPerRound, len(teams))
		fmt.Printf("Round: %v\n", round)

		var currentGames []models.Game

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
				currentGames = append(currentGames, models.Game{Team1: team1, Team2: team2})
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
}
func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, err := models.EventsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var results []models.Event
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No events found"}
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventId, err := primitive.ObjectIDFromHex(r.PathValue("event_id"))
	if err != nil {
		panic(err)
	}

	var result models.Event
	err = models.EventsCollection.FindOne(context.TODO(), bson.M{"_id": eventId}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			errorResponse := map[string]string{"error": "Event not found"}
			json.NewEncoder(w).Encode(errorResponse)
			return
		}
		panic(err)
	}
	fmt.Println(result)
	json.NewEncoder(w).Encode(result.Teams)
}
func main() {
	configs.ConnectDB()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/events", eventsHandler)
	mux.HandleFunc("GET /api/{event_id}/teams", teamsHandler)
	// mux.HandleFunc("GET /api/{event_id}/pools", poolsHandler)
	// mux.HandleFunc("GET /api/{event_id}/seeding", seedingHandler)
	// mux.HandleFunc("GET /api/{event_id}/bracket", bracketHandler)
	// mux.HanleFunc("GET /api/{event_id}/results", resultsHandler)

	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", mux)
}
