package main

import (
	"fmt"
	"net/http"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"github.com/CrossoversForCures/Tournament-Scoring/routes"
)

func main() {
	configs.ConnectDB()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/start", routes.StartHandler)
	mux.HandleFunc("GET /api/home", routes.HomeHandler)
	mux.HandleFunc("GET /api/{event_id}/teams", routes.TeamsHandler)
	mux.HandleFunc("GET /api/{event_id}/pools", routes.PoolsHandler)
	// mux.HandleFunc("GET /api/{event_id}/seeding", seedingHandler)
	// mux.HandleFunc("GET /api/{event_id}/bracket", bracketHandler)
	// mux.HanleFunc("GET /api/{event_id}/results", resultsHandler)

	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", mux)
}
