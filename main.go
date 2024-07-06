package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/CrossoversForCures/Tournament-Scoring/models"
)

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
	event_id := r.PathValue("event_id")
	var result models.Event
	err := models.EventsCollection.FindOne(context.TODO(), bson.D{{Key: "ID", Value: event_id}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			w.WriteHeader(http.StatusNotFound)
			errorResponse := map[string]string{"error": "Event not found"}
			json.NewEncoder(w).Encode(errorResponse)
		}
		panic(err)
	}

	json.NewEncoder(w).Encode(result.Teams)
}
func main() {
	models.ConnectDB()
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
