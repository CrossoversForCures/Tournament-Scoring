package stripeFunctions

import (
	"fmt"
	"log"

	"github.com/CrossoversForCures/Tournament-Scoring/models"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
)

func GetAllTeams() ([]*models.Team, error) {
	// Initialize params for listing checkout sessions
	params := &stripe.CheckoutSessionListParams{}

	// Create a slice to store all sessions
	var teamRegistrations []*models.Team

	// Use an iterator to paginate through all checkout sessions
	i := session.List(params)
	for i.Next() {
		s := i.CheckoutSession()
		if s.CustomFields != nil && len(s.CustomFields) > 2 {
			team, err := processCheckoutSessionTest(s)
			if err != nil {
				continue
			}
			teamRegistrations = append(teamRegistrations, &team)
		}

	}
	// Handle any error encountered during the iteration
	if err := i.Err(); err != nil {
		log.Fatalf("Error listing sessions: %v", err)
	}

	fmt.Printf("Total sessions retrieved: %d\n", len(teamRegistrations))
	return teamRegistrations, nil
}

func processCheckoutSessionTest(s *stripe.CheckoutSession) (models.Team, error) {
	team := models.Team{}

	for _, field := range s.CustomFields {
		switch field.Key {
		case "teamname":
			if field.Text != nil {
				team.Name = field.Text.Value
			}
		case "division":
			if field.Dropdown != nil {
				team.Division = field.Dropdown.Value
			}
		case "numberofplayers":
			if field.Dropdown != nil {
				team.NumPlayers = field.Dropdown.Value
			}
		}
	}
	if team.Name == "" || team.Division == "" || team.NumPlayers == "" {
		return models.Team{}, fmt.Errorf("incomplete team information: TeamName: %s, Division: %s, NumPlayers: %s",
			team.Name, team.Division, team.NumPlayers)
	}
	return team, nil
}
