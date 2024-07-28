package main

import (
	"fmt"
	"net/http"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"github.com/CrossoversForCures/Tournament-Scoring/backend/routes"
	"github.com/rs/cors"

	"github.com/stripe/stripe-go/v79"
)

func main() {
	configs.ConnectDB()
	stripe.Key = configs.GetStripeKey()
	// models.InitiateTestTeams()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/home", routes.HomeHandler)
	mux.HandleFunc("GET /api/{event_slug}/teams", routes.TeamsHandler)
	mux.HandleFunc("GET /api/{event_slug}/pools", routes.PoolsHandler)
	mux.HandleFunc("GET /api/{event_slug}/seeding", routes.SeedingHandler)
	mux.HandleFunc("GET /api/{event_slug}/bracket", routes.BracketHandler)
	mux.HandleFunc("GET /api/{event_slug}/results", routes.ResultsHandler)

	mux.HandleFunc("POST /api/{event_slug}/start-pools", routes.StartPoolsHandler)
	mux.HandleFunc("POST /api/{event_slug}/start-elimination", routes.StartEliminationHandler)

	mux.HandleFunc("POST /api/update-pool", routes.UpdatePoolsHandler)
	mux.HandleFunc("POST /api/update-elimination", routes.UpdateElimHandler)
	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	handler := corsOptions.Handler(mux)
	fmt.Println("Starting server on port 8000")
	http.ListenAndServe(":8000", handler)
}
