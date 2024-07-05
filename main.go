package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/CrossoversForCures/Tournament-Scoring/models"
)

var err error

func eventsHandler(w http.ResponseWriter, r *http.Request) {
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
		response, err := json.Marshal(results)
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(response)
	}

}

func main() {
	models.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/events", eventsHandler)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", mux)
}
