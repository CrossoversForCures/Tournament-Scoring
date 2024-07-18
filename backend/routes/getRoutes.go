package routes

import (
	"cmp"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/models"
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

	eventSlug := r.PathValue("event_slug")

	results := models.GetTeams(eventSlug)

	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No teams found for event"}
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}

func PoolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventSlug := r.PathValue("event_slug")

	event := models.GetEvent(eventSlug)

	if event.Status < 1 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "Pool round hasn't started yet"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	results := models.GetPools(eventSlug)

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

		json.NewEncoder(w).Encode(games)
	}
}

func SeedingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventSlug := r.PathValue("event_slug")

	event := models.GetEvent(eventSlug)

	if event.Status < 2 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "Elimination round hasn't started yet"}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	results := models.GetTeams(eventSlug)

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

	json.NewEncoder(w).Encode(results)
}

func ElimHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	eventSlug := r.PathValue("event_slug")

	bracket := models.GetBracket(eventSlug)

	json.NewEncoder(w).Encode(bracket)
}
