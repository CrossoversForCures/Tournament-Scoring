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
	mux.HandleFunc("GET /api/{event_name}/teams", routes.TeamsHandler)
	mux.HandleFunc("GET /api/{event_name}/pools", routes.PoolsHandler)
	mux.HandleFunc("GET /api/{event_name}/seeding", routes.SeedingHandler)
	mux.HandleFunc("GET /api/{event_name}/bracket", routes.ElimHandler)

	mux.HandleFunc("POST /api/{event_name}/start-pools", routes.StartPoolsHandler)
	mux.HandleFunc("POST /api/{event_name}/start-elimination", routes.StartEliminationHandler)

	mux.HandleFunc("POST /api/{event_name}/update-pool", routes.UpdatePoolsHandler)
	mux.HandleFunc("POST /api/{event_name}/update-elim", routes.UpdateElimHandler)

	// mux.HanleFunc("GET /api/{eventId}/results", resultsHandler)

	handler := cors.Default().Handler(mux)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", handler)
}
