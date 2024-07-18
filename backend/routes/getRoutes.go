package routes

import (
	"cmp"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type response struct {
		Year   int            `json:"year"`
		Events []models.Event `json:"events"`
	}

	results := models.GetEvents()

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No events found"}
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		newResponse := response{
			Year:   time.Now().Year(),
			Events: results,
		}
		json.NewEncoder(w).Encode(newResponse)
	}
}

func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type response struct {
		Teams []models.Team `json:"teams"`
	}

	eventId, err := primitive.ObjectIDFromHex(r.PathValue("event_id"))
	if err != nil {
		panic(err)
	}

	results := models.GetTeams(eventId)

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No teams found for event"}
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		newResponse := response{
			Teams: results,
		}
		json.NewEncoder(w).Encode(newResponse)
	}
}

func PoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type response struct {
		Games map[string][]models.PoolGame `json:"games"`
	}

	eventId, err := primitive.ObjectIDFromHex(r.PathValue("event_id"))
	if err != nil {
		panic(err)
	}

	event := models.GetEvent(eventId)
	if event.Status < 1 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "Pool round hasn't started yet"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	results := models.GetPools(eventId)

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No pool games found"}
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		games := make(map[string][]models.PoolGame)
		for _, game := range results {
			key := fmt.Sprintf("%d", game.Round)
			games[key] = append(games[key], game)
		}

		newResponse := response{
			Games: games,
		}

		json.NewEncoder(w).Encode(newResponse)
	}
}

func SeedingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type response struct {
		Seeding []models.Team `json:"seeding"`
	}

	eventId, err := primitive.ObjectIDFromHex(r.PathValue("event_id"))
	if err != nil {
		panic(err)
	}

	event := models.GetEvent(eventId)
	if event.Status < 2 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "Elimination round hasn't started yet"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	results := models.GetTeams(eventId)

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No teams found"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	slices.SortFunc(results, func(a, b models.Team) int {
		return cmp.Or(
			cmp.Compare(a.Seeding, b.Seeding),
		)
	})

	newResponse := response{
		Seeding: results,
	}
	json.NewEncoder(w).Encode(newResponse)
}

func ElimHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventId, err := primitive.ObjectIDFromHex(r.PathValue("event_id"))
	if err != nil {
		panic(err)
	}

	currentBracket := models.GetEvent(eventId).ElimBracket
	bracketJSON := currentBracket.ToDisplayNode(currentBracket.Root)
	json.NewEncoder(w).Encode(bracketJSON)
}
