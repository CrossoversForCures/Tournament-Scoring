package models

import (
	"cmp"
	"context"
	"fmt"
	"math"
	"slices"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BracketNode struct {
	Team    string             `bson:"team,omitempty" json:"team,omitempty"`
	TeamID  primitive.ObjectID `bson:"teamId,omitempty" json:"teamId,omitempty"`
	Seeding int                `bson:"seeding,omitempty" json:"seeding,omitempty"`
	Left    *BracketNode       `bson:"left,omitempty" json:"left,omitempty"`
	Right   *BracketNode       `bson:"right,omitempty" json:"right,omitempty"`
}

type Bracket struct {
	Event string       `bson:"event,omitempty" json:"event,omitempty"`
	Root  *BracketNode `bson:"root,omitempty" json:"root,omitempty"`
}

func GetBracket(eventSlug string) Bracket {
	var result Bracket
	err := configs.BracketsCollection.FindOne(context.TODO(), bson.D{{Key: "event", Value: eventSlug}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func InsertBracket(newBracket Bracket) {
	_, err := configs.BracketsCollection.InsertOne(context.TODO(), newBracket)
	if err != nil {
		panic(err)
	}
}

func UpdateBracket(eventSlug string, update bson.D) {
	_, err := configs.BracketsCollection.UpdateOne(context.TODO(), bson.D{{Key: "event", Value: eventSlug}}, update)
	if err != nil {
		panic(err)
	}
}

func SeedTeams(eventSlug string) {
	results := GetTeams(eventSlug)

	slices.SortFunc(results, func(a, b Team) int {
		return cmp.Or(
			cmp.Compare(b.PoolsWon, a.PoolsWon),
			cmp.Compare(b.TotalPoints, a.TotalPoints),
		)
	})

	for i := 0; i < len(results); i++ {
		UpdateTeam(results[i].ID, bson.D{{Key: "$set", Value: bson.D{{Key: "seeding", Value: i + 1}}}})
	}
}

func MakeBracket(eventSlug string) {
	_, err := configs.BracketsCollection.DeleteOne(context.TODO(), bson.D{{Key: "event", Value: eventSlug}})
	if err != nil {
		panic(err)
	}
	teams := GetTeams(eventSlug)
	matchups := getBracket(teams)

	bracket := buildBracketTree(matchups, eventSlug)
	resolveByes(bracket.Root)
	PrintBracketTree(bracket.Root, 0)
	InsertBracket(*bracket)
}

func getBracket(teams []Team) [][]*Team {
	participantsCount := len(teams)
	rounds := int(math.Ceil(math.Log2(float64(participantsCount))))
	// bracketSize := int(math.Pow(2, float64(rounds)))
	// requiredByes := bracketSize - participantsCount

	// fmt.Printf("Number of participants: %d\n", participantsCount)
	// fmt.Printf("Number of rounds: %d\n", rounds)
	// fmt.Printf("Bracket size: %d\n", bracketSize)
	// fmt.Printf("Required number of byes: %d\n", requiredByes)

	if participantsCount < 2 {
		return [][]*Team{}
	}

	matches := [][]*Team{{&teams[0], &teams[1]}}

	for round := 1; round < rounds; round++ {
		roundMatches := [][]*Team{}
		sum := int(math.Pow(2, float64(round+1))) + 1

		for i := 0; i < len(matches); i++ {
			home := changeIntoBye(matches[i][0].Seeding, participantsCount, teams)
			away := changeIntoBye(sum-matches[i][0].Seeding, participantsCount, teams)
			roundMatches = append(roundMatches, []*Team{home, away})

			home = changeIntoBye(sum-matches[i][1].Seeding, participantsCount, teams)
			away = changeIntoBye(matches[i][1].Seeding, participantsCount, teams)
			roundMatches = append(roundMatches, []*Team{home, away})
		}
		matches = roundMatches
	}

	return matches
}

func changeIntoBye(seed, participantsCount int, teams []Team) *Team {
	if seed <= participantsCount {
		return &teams[seed-1]
	}
	return nil
}

func buildBracketTree(matchups [][]*Team, eventSlug string) *Bracket {
	if len(matchups) == 0 {
		return nil
	}

	nodes := make([]*BracketNode, 0, len(matchups)*2)

	// Create leaf nodes
	for _, matchup := range matchups {
		for _, team := range matchup {
			if team == nil {
				nodes = append(nodes, &BracketNode{
					Team: "BYE",
				}) // Represent a bye
			} else {
				nodes = append(nodes, &BracketNode{
					Team:    team.Name,
					TeamID:  team.ID,
					Seeding: team.Seeding,
				})
			}
		}
	}

	// Build the tree bottom-up
	for len(nodes) > 1 {
		var parents []*BracketNode
		for i := 0; i < len(nodes); i += 2 {
			parent := &BracketNode{}
			if i+1 < len(nodes) {
				parent.Left = nodes[i]
				parent.Right = nodes[i+1]
			} else {
				parent.Left = nodes[i]
			}
			parents = append(parents, parent)
		}
		nodes = parents
	}

	newBracket := Bracket{
		Event: eventSlug,
		Root:  nodes[0],
	}
	return &newBracket
}

func resolveByes(root *BracketNode) {
	if root == nil {
		return
	}

	// Check if this node is a matchup (has children)
	if root.Left != nil || root.Right != nil {
		if root.Left.Team != "BYE" && root.Right.Team == "BYE" {
			bypassBye(root, root.Left)
		} else if root.Right.Team != "BYE" && root.Left.Team == "BYE" {
			bypassBye(root, root.Right)
		} else {
			resolveByes(root.Left)
			resolveByes(root.Right)
		}
	}
}

func bypassBye(parent *BracketNode, winner *BracketNode) {
	parent.Team = winner.Team
	parent.TeamID = winner.TeamID
	parent.Seeding = winner.Seeding
}

func SetWinner(root *BracketNode, teamID primitive.ObjectID) bool {
	if root == nil {
		return false
	}

	// Check if either child is the team we're looking for
	if root.Left != nil && root.Left.TeamID == teamID {
		root.Team = root.Left.Team
		root.TeamID = root.Left.TeamID
		root.Seeding = root.Left.Seeding
		return true
	}
	if root.Right != nil && root.Right.TeamID == teamID {
		root.Team = root.Right.Team
		root.TeamID = root.Right.TeamID
		root.Seeding = root.Right.Seeding
		return true
	}

	// If not found in immediate children, recursively search left and right subtrees
	if SetWinner(root.Left, teamID) {
		return true
	}
	if SetWinner(root.Right, teamID) {
		return true
	}

	return false
}

func PrintBracketTree(node *BracketNode, depth int) {
	if node == nil {
		return
	}

	PrintBracketTree(node.Right, depth+1)

	for i := 0; i < depth; i++ {
		fmt.Print("    ")
	}
	if node.Team == "" {
		fmt.Println("[-]")
	} else {
		fmt.Printf("[%s (Seed %d)]\n", node.Team, node.Seeding)
	}

	PrintBracketTree(node.Left, depth+1)
}
