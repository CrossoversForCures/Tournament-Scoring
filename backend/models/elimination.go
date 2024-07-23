package models

import (
	"cmp"
	"context"
	"fmt"
	"math"
	"slices"
	"sort"

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
	_, err := configs.BracketsCollection.DeleteMany(context.TODO(), bson.D{{Key: "event", Value: eventSlug}})
	if err != nil {
		panic(err)
	}
	teams := GetTeams(eventSlug)
	slices.SortFunc(teams, func(a, b Team) int {
		return cmp.Or(
			cmp.Compare(a.Seeding, b.Seeding),
		)
	})
	matchups := getMatchups(teams)
	bracket := buildTree(eventSlug, matchups)
	resolveByes(bracket.Root)
	assignCourts(bracket)
	InsertBracket(*bracket)
}

func getMatchups(teams []Team) [][]Team {
	numTeams := len(teams)
	rounds := int(math.Ceil(math.Log2(float64(numTeams))))

	matches := [][]Team{{teams[0], teams[1]}}

	for round := 1; round < rounds; round++ {
		var roundMatches [][]Team
		sum := int(math.Pow(2, float64(round+1))) + 1

		for i := 0; i < len(matches); i++ {
			var home Team
			var away Team

			if matches[i][0].Seeding <= numTeams {
				home = matches[i][0]
			} else {
				home = Team{}
			}
			if sum-matches[i][0].Seeding <= numTeams {
				away = teams[sum-matches[i][0].Seeding-1]
			} else {
				away = Team{}
			}
			roundMatches = append(roundMatches, []Team{home, away})
			if sum-matches[i][1].Seeding <= numTeams {
				home = teams[sum-matches[i][1].Seeding-1]
			} else {
				home = Team{}
			}
			if matches[i][1].Seeding <= numTeams {
				away = matches[i][1]
			} else {
				away = Team{}
			}
			roundMatches = append(roundMatches, []Team{home, away})
		}
		matches = roundMatches
	}

	return matches
}

func buildTree(eventSlug string, matchups [][]Team) *Bracket {
	if len(matchups) == 0 {
		return nil
	}

	nodes := make([]*BracketNode, 0, len(matchups)*2)

	// Create leaf nodes
	for _, matchup := range matchups {
		for _, team := range matchup {
			if team.Name == "" {
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
		root.Court = "-" + root.Court
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
		root.Court = "-" + root.Court
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

// Helper
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

func RankTeams(eventSlug string) {
	bracket := GetBracket(eventSlug)
	teams := orderTree(bracket.Root)
	sortTeams(teams, bracket.Rounds)
}

// Helper
func orderTree(root *BracketNode) []*BracketNode {
	if root == nil {
		return nil
	}

	// Slice to store unique teams in order
	var result []*BracketNode

	// Use a map to keep track of processed teams
	processed := make(map[primitive.ObjectID]bool)

	// Queue for BFS
	queue := []*BracketNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Process current node if not processed
			if node.TeamID != primitive.NilObjectID && !processed[node.TeamID] {
				result = append(result, node)
				processed[node.TeamID] = true
			}

			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

// Helper
func sortTeams(teams []*BracketNode, rounds int) {
	if rounds <= 2 || len(teams) == 0 {
		return // No sorting needed
	}

	for round := 1; round < rounds; round++ {
		if round == rounds-1 {
			sortSection(teams, int(math.Pow(2, float64(round))), len(teams)-1)
		} else {
			sortSection(teams, int(math.Pow(2, float64(round))), int(math.Pow(2, float64(round+1))))
		}
	}

	for i := 0; i < len(teams); i++ {
		if i == 3 {
			UpdateTeam(teams[i].TeamID, getUpdate(3))
		} else {
			UpdateTeam(teams[i].TeamID, getUpdate(i+1))
		}
	}
}

// Helper
func getUpdate(rank int) bson.D {
	return bson.D{{Key: "$set", Value: bson.D{{Key: "rank", Value: rank}}}}
}

// Helper
func sortSection(nodes []*BracketNode, start, end int) {
	// Check if the indices are valid
	if start < 0 || end >= len(nodes) || start >= end {
		return
	}

	// Sort the section directly
	sort.Slice(nodes[start:end+1], func(i, j int) bool {
		return nodes[start+i].Seeding < nodes[start+j].Seeding
	})
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
