package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"github.com/CrossoversForCures/Tournament-Scoring/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StartHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cursor, err := configs.EventsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	var events []models.Event
	if err = cursor.All(context.TODO(), &events); err != nil {
		panic(err)
	}

	for _, event := range events {
		eventId, _ := primitive.ObjectIDFromHex(event.ID.Hex())
		models.SortPools(eventId)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Year   int
		Events []models.Event
	}

	w.Header().Set("Content-Type", "application/json")

	cursor, err := configs.EventsCollection.Find(context.TODO(), bson.D{})
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
		newResponse := response{
			Year:   time.Now().Year(),
			Events: results,
		}
		json.NewEncoder(w).Encode(newResponse)
	}
}

func TeamsHandler(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Teams []models.Team
	}

	w.Header().Set("Content-Type", "application/json")
	eventId, err := primitive.ObjectIDFromHex(r.PathValue("event_id"))
	if err != nil {
		panic(err)
	}

	cursor, err := configs.TeamsCollection.Find(context.TODO(), bson.M{"event_id": eventId})
	var results []models.Team
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
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
	eventId, err := primitive.ObjectIDFromHex(r.PathValue("event_id"))
	if err != nil {
		panic(err)
	}

	cursor, err := configs.GamesCollection.Find(context.TODO(), bson.M{"event_id": eventId})
	var results []models.PoolGame
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if len(results) == 0 {
		w.WriteHeader(http.StatusNotFound)
		errorResponse := map[string]string{"error": "No pool games found"}
		json.NewEncoder(w).Encode(errorResponse)
	} else {
		json.NewEncoder(w).Encode(results)
	}
}
