package main

import (
	"fmt"
	"log"

	"github.com/CrossoversForCures/Tournament-Scoring/configs"
	"github.com/CrossoversForCures/Tournament-Scoring/stripeFunctions"
	"github.com/stripe/stripe-go/v79"
)

func init() {
	stripe.Key = configs.GetStripeKey()
}

func main() {
	// Call the getTeams function
	teams, err := stripeFunctions.GetAllTeams()
	if err != nil {
		log.Fatalf("Error fetching teams: %v", err)
	}

	// Log the results
	for _, team := range teams {
		fmt.Printf("Team: %+v\n", team)
	}

}
