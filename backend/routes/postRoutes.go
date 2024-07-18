package routes

import (
	"encoding/json"
	"net/http"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartPoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventSlug := r.PathValue("event_slug")

	models.SortPools(eventSlug)
	models.UpdateEvent(eventSlug, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: 1}}}})

	response := map[string]string{"response": "Pools successfuly started"}
	json.NewEncoder(w).Encode(response)
}

func StartEliminationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventSlug := r.PathValue("event_slug")

	models.SeedTeams(eventSlug)
	models.MakeBracket(eventSlug)
	models.UpdateEvent(eventSlug, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: 2}}}})

	response := map[string]string{"response": "Elimination successfuly started"}
	json.NewEncoder(w).Encode(response)
}

func UpdatePoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type request struct {
		Team1Score int `bson:"team1_score" json:"team1Score"`
		Team2Score int `bson:"team2_score" json:"team2Score"`
	}

	gameId, err := primitive.ObjectIDFromHex(r.PathValue("game_id"))
	if err != nil {
		panic(err)
	}

	var newRequest request
	err = json.NewDecoder(r.Body).Decode(&newRequest)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "team1Score", Value: newRequest.Team1Score}, {Key: "team2Score", Value: newRequest.Team2Score}}}}
	models.UpdatePool(gameId, update)

	poolGame := models.GetPool(gameId)
	team1 := models.GetTeam(poolGame.Team1)
	team2 := models.GetTeam(poolGame.Team2)
	team1Update := bson.D{{Key: "$set", Value: bson.D{{Key: "totalPoints", Value: team1.TotalPoints + newRequest.Team1Score}}}}
	team2Update := bson.D{{Key: "$set", Value: bson.D{{Key: "totalPoints", Value: team2.TotalPoints + newRequest.Team2Score}}}}

	models.UpdateTeam(team1.ID, team1Update)
	models.UpdateTeam(team2.ID, team2Update)

	var winner models.Team
	if newRequest.Team1Score > newRequest.Team2Score {
		winner = team1
	} else {
		winner = team2
	}
	models.UpdateTeam(winner.ID, bson.D{{Key: "$set", Value: bson.D{{Key: "poolsWon", Value: winner.PoolsWon + 1}}}})

	response := map[string]string{"response": "Game successfuly updated"}
	json.NewEncoder(w).Encode(response)
}

func UpdateElimHandler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	// teamId, err := primitive.ObjectIDFromHex(r.PathValue("team_id"))
	// if err != nil {
	// 	panic(err)
	// }

	// // models.ElimBracket.SetWinner(models.GetTeam(teamId))
	// // models.ElimBracket.PrintTree(models.ElimBracket.Root, 0)
}
