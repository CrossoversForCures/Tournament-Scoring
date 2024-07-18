package main

import (
	"fmt"
	"net/http"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"github.com/CrossoversForCures/Tournament-Scoring/backend/routes"
	"github.com/rs/cors"
)

func main() {
	configs.ConnectDB()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/home", routes.HomeHandler)
	mux.HandleFunc("GET /api/{event_id}/teams", routes.TeamsHandler)
	mux.HandleFunc("GET /api/{event_id}/pools", routes.PoolsHandler)
	mux.HandleFunc("GET /api/{event_id}/seeding", routes.SeedingHandler)
	mux.HandleFunc("GET /api/{eventId}/bracket", routes.ElimHandler)

	mux.HandleFunc("POST /api/{event_id}/start-pools", routes.StartPoolsHandler)
	mux.HandleFunc("POST /api/{event_id}/start-elimination", routes.StartEliminationHandler)

	mux.HandleFunc("POST /api/{game_id}/score", routes.UpdatePoolsHandler)
	mux.HandleFunc("POST /api/{team_id}/winner", routes.UpdateElimHandler)

	// mux.HanleFunc("GET /api/{eventId}/results", resultsHandler)

	handler := cors.Default().Handler(mux)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", handler)
}
