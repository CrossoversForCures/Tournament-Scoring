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
		GameID     primitive.ObjectID `bson:"gameId" json:"gameId"`
		Team1Score json.Number        `bson:"team1Score" json:"team1Score"`
		Team2Score json.Number        `bson:"team2Score" json:"team2Score"`
	}

	var newRequest request
	err := json.NewDecoder(r.Body).Decode(&newRequest)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	gameId, err := primitive.ObjectIDFromHex(newRequest.GameID.Hex())
	if err != nil {
		panic(err)
	}
	team1Score, _ := newRequest.Team1Score.Int64()
	team2Score, _ := newRequest.Team2Score.Int64()

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "team1Score", Value: newRequest.Team1Score}, {Key: "team2Score", Value: newRequest.Team2Score}}}}
	models.UpdatePool(gameId, update)

	poolGame := models.GetPool(gameId)
	team1 := models.GetTeam(poolGame.Team1)
	team2 := models.GetTeam(poolGame.Team2)
	team1Update := bson.D{{Key: "$set", Value: bson.D{{Key: "totalPoints", Value: team1.TotalPoints + int(team1Score)}}}}
	team2Update := bson.D{{Key: "$set", Value: bson.D{{Key: "totalPoints", Value: team2.TotalPoints + int(team2Score)}}}}

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
	w.Header().Set("Content-Type", "application/json")

	type request struct {
		TeamID primitive.ObjectID `bson:"teamId" json:"teamId"`
	}

	var newRequest request
	err := json.NewDecoder(r.Body).Decode(&newRequest)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	teamId, err := primitive.ObjectIDFromHex(newRequest.TeamID.Hex())
	if err != nil {
		panic(err)
	}

	team := models.GetTeam(teamId)
	bracket := models.GetBracket(team.Event)
	if models.SetWinner(bracket.Root, team.ID, &bracket.Courts) {
		models.UpdateBracket(team.Event, bson.D{{Key: "$set", Value: bson.D{{Key: "root", Value: bracket.Root}, {Key: "courts", Value: bracket.Courts}}}})
	} else {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No valid matchup"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	if bracket.Root.Team != "" {
		models.RankTeams(team.Event)
		models.UpdateEvent(team.Event, bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: 3}}}})
	}

}
