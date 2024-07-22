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
	Court   string             `bson:"court,omitempty" json:"court,omitempty"`
}

type Bracket struct {
	Event  string       `bson:"event,omitempty" json:"event,omitempty"`
	Rounds int          `bson:"rounds,omitempty" json:"rounds,omitempty"`
	Courts []string     `bson:"courts,omitempty" json:"courts,"`
	Root   *BracketNode `bson:"root,omitempty" json:"root,omitempty"`
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
	teams := GetTeams(eventSlug)
	matchups := getMatchups(teams)

	bracket := buildTree(eventSlug, matchups)
	resolveByes(bracket.Root)
	assignCourts(bracket)
	InsertBracket(*bracket)
}

func getMatchups(teams []Team) [][]*Team {
	participantsCount := len(teams)
	rounds := int(math.Ceil(math.Log2(float64(participantsCount))))

	if participantsCount < 2 {
		return [][]*Team{}
	}

	matches := [][]Team{{teams[0], teams[1]}}

	for round := 1; round < rounds; round++ {
		roundMatches := [][]Team{}
		sum := int(math.Pow(2, float64(round+1))) + 1

		for i := 0; i < len(matches); i++ {
			home := changeIntoBye(matches[i][0].Seeding, participantsCount, teams)
			away := changeIntoBye(sum-matches[i][0].Seeding, participantsCount, teams)
			roundMatches = append(roundMatches, []Team{home, away})

			home = changeIntoBye(sum-matches[i][1].Seeding, participantsCount, teams)
			away = changeIntoBye(matches[i][1].Seeding, participantsCount, teams)
			roundMatches = append(roundMatches, []Team{home, away})
		}
		matches = roundMatches
	}

	result := make([][]*Team, len(matches))
	for i, match := range matches {
		result[i] = make([]*Team, len(match))
		for j := range match {
			result[i][j] = &match[j]
		}
	}
	return result
}

// Helper
func changeIntoBye(seed, participantsCount int, teams []Team) Team {
	if seed <= 0 || seed > participantsCount {
		return Team{} // Return an empty Team for byes or invalid seeds
	}
	return teams[seed-1]
}

func buildTree(eventSlug string, matchups [][]*Team) *Bracket {
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
		Event:  eventSlug,
		Root:   nodes[0],
		Rounds: int(math.Log2(float64(len(matchups) * 2))),
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

// Helper
func bypassBye(parent *BracketNode, winner *BracketNode) {
	parent.Team = winner.Team
	parent.TeamID = winner.TeamID
	parent.Seeding = winner.Seeding
}

func assignCourts(bracket *Bracket) {
	courts := []string{"A", "B", "C", "D"}
	assignCourtsAtLayer(bracket.Root, bracket.Rounds, 1, &courts)
	assignCourtsAtLayer(bracket.Root, bracket.Rounds-1, 1, &courts)
	bracket.Courts = courts
}

// Helper
func assignCourtsAtLayer(node *BracketNode, targetDepth int, currentDepth int, courts *[]string) {
	if node == nil {
		return
	}

	if currentDepth == targetDepth {
		if node.Left != nil && node.Right != nil &&
			node.Left.Team != "" && node.Right.Team != "" &&
			node.Left.Team != "BYE" && node.Right.Team != "BYE" {
			if len(*courts) > 0 {
				node.Court = (*courts)[0]
				*courts = (*courts)[1:]
			}
		} else if (node.Left != nil && node.Left.Team == "BYE") ||
			(node.Right != nil && node.Right.Team == "BYE") {
			node.Court = "N/A"
		}
		return
	}

	assignCourtsAtLayer(node.Left, targetDepth, currentDepth+1, courts)
	assignCourtsAtLayer(node.Right, targetDepth, currentDepth+1, courts)
}

func SetWinner(root *BracketNode, teamID primitive.ObjectID, courts *[]string) bool {
	if setWinnerRecursive(root, teamID, courts) {
		if assignNext(root, (*courts)[0]) {
			*courts = (*courts)[1:]
		}
		return true
	}
	return false
}

// Helper
func setWinnerRecursive(root *BracketNode, teamID primitive.ObjectID, courts *[]string) bool {
	if root == nil {
		return false
	}

	// Check if either child is the team we're looking for
	if root.Right != nil && root.Left.TeamID == teamID {
		if root.Right.Team == "" {
			return false
		}
		root.Team = root.Left.Team
		root.TeamID = root.Left.TeamID
		root.Seeding = root.Left.Seeding
		*courts = append(*courts, root.Court)
		return true
	}
	if root.Left != nil && root.Right.TeamID == teamID {
		if root.Left.Team == "" {
			return false
		}
		root.Team = root.Right.Team
		root.TeamID = root.Right.TeamID
		root.Seeding = root.Right.Seeding
		*courts = append(*courts, root.Court)
		return true
	}

	// If not found in immediate children, recursively search left and right subtrees
	if setWinnerRecursive(root.Left, teamID, courts) {
		return true
	}
	if setWinnerRecursive(root.Right, teamID, courts) {
		return true
	}

	return false
}

func assignNext(node *BracketNode, court string) bool {
	if node == nil {
		return false
	}

	// Check if this node is a valid, unassigned matchup
	if node.Left != nil && node.Right != nil &&
		node.Left.Team != "" && node.Right.Team != "" &&
		node.Left.Team != "BYE" && node.Right.Team != "BYE" &&
		node.Court == "" {
		node.Court = court
		return true
	}

	// Recursively check left subtree
	if assignNext(node.Left, court) {
		return true
	}

	// Recursively check right subtree
	return assignNext(node.Right, court)
}

// Helper function for debugging
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
