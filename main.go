package main

import (
	"fmt"
	"net/http"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"github.com/CrossoversForCures/Tournament-Scoring/routes"

	"github.com/rs/cors"
)

func main() {
	configs.ConnectDB()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/home", routes.HomeHandler)
	mux.HandleFunc("GET /api/{event_id}/teams", routes.TeamsHandler)
	mux.HandleFunc("GET /api/{event_id}/pools", routes.PoolsHandler)
	mux.HandleFunc("GET /api/{event_id}/seeding", routes.SeedingHandler)

	mux.HandleFunc("POST /api/{event_id}/start-pools", routes.StartPoolsHandler)
	mux.HandleFunc("POST /api/{event_id}/start-elimination", routes.StartEliminationHandler)
	mux.HandleFunc("POST /api/{game_id}/score", routes.UpdatePoolsHandler)

	// mux.HandleFunc("GET /api/{eventId}/bracket", bracketHandler)
	// mux.HanleFunc("GET /api/{eventId}/results", resultsHandler)

	// mux.HandleFunc("POST /api/poolgame/{game_id}", routes.UpdatePoolsHandler)
	handler := cors.Default().Handler(mux)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", handler)
}
