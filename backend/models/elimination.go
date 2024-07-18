package models

import (
	"cmp"
	"context"
	"fmt"
	"slices"

	"github.com/CrossoversForCures/Tournament-Scoring/backend/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeedTeams(eventId primitive.ObjectID) {
	results := GetTeams(eventId)
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

func MakeBracket(eventId primitive.ObjectID) {
	_, err := configs.ElimGamesCollection.DeleteMany(context.TODO(), bson.D{{Key: "eventId", Value: eventId}})
	if err != nil {
		panic(err)
	}
	teams := GetTeams(eventId)
	slices.SortFunc(teams, func(a, b Team) int {
		return cmp.Or(
			cmp.Compare(a.Seeding, b.Seeding),
		)
	})

	roundsList := [6]int{2, 4, 8, 16, 32, 64}
	var bracket int
	for _, r := range roundsList {
		if len(teams) <= r {
			bracket = r
			break
		}
	}

	teams = append(teams, make([]Team, bracket-len(teams))...)
	newBracket := *NewBracket(teams)
	newBracket.ProcessByes()

	UpdateEvent(eventId, bson.D{{Key: "$set", Value: bson.D{{Key: "elimBracket", Value: newBracket}}}})
	newBracket.PrintTree(newBracket.Root, 0)
}

type BracketNode struct {
	Team   *Team
	Left   *BracketNode
	Right  *BracketNode
	Parent *BracketNode
}

type Bracket struct {
	Root *BracketNode
}

func NewBracket(teams []Team) *Bracket {
	if len(teams) == 0 {
		return &Bracket{}
	}

	leaves := make([]*BracketNode, len(teams))
	for i := range teams {
		leaves[i] = &BracketNode{Team: &teams[i]}
	}

	// Build the tree bottom-up
	for len(leaves) > 1 {
		var parents []*BracketNode
		for i := 0; i < len(leaves); i += 2 {
			parent := &BracketNode{}
			parent.Left = leaves[i]
			leaves[i].Parent = parent
			if i+1 < len(leaves) {
				parent.Right = leaves[i+1]
				leaves[i+1].Parent = parent
			}
			parents = append(parents, parent)
		}
		leaves = parents
	}

	return &Bracket{Root: leaves[0]}
}

func (bt *Bracket) SetWinner(team Team) {
	node := bt.findNode(bt.Root, team)
	if node == nil || node.Parent == nil {
		return
	}

	parent := node.Parent
	parent.Team = node.Team
}

// ProcessByes goes through the tree and moves up teams paired with "BYE"
func (bt *Bracket) ProcessByes() {
	bt.processByesRecursive(bt.Root)
}

// Recursive helper function to process byes
func (bt *Bracket) processByesRecursive(node *BracketNode) {
	if node == nil {
		return
	}

	// If this is a parent of leaf nodes
	if (node.Left != nil && node.Left.Team != nil) && (node.Right != nil && node.Right.Team != nil) {
		if node.Left.Team.Name == "BYE" {
			bt.SetWinner(*node.Right.Team)
		} else if node.Right.Team.Name == "BYE" {
			bt.SetWinner(*node.Left.Team)
		}
		return // No need to process further down this branch
	}

	// Continue searching in left and right subtrees
	bt.processByesRecursive(node.Left)
	bt.processByesRecursive(node.Right)
}

// Helper function to find a node with a specific team
func (bt *Bracket) findNode(node *BracketNode, team Team) *BracketNode {
	if node == nil {
		return nil
	}
	if node.Team != nil && node.Team.ID == team.ID {
		return node
	}
	left := bt.findNode(node.Left, team)
	if left != nil {
		return left
	}
	return bt.findNode(node.Right, team)
}

// PrintTree is a helper function to print the tree (for debugging)
func (bt *Bracket) PrintTree(node *BracketNode, level int) {
	if node == nil {
		return
	}
	for i := 0; i < level; i++ {
		fmt.Print("  ")
	}
	if node.Team != nil {
		fmt.Printf("Team: %s\n", node.Team.Name)
	} else {
		fmt.Println("Empty")
	}
	bt.PrintTree(node.Left, level+1)
	bt.PrintTree(node.Right, level+1)
}

type DisplayNode struct {
	Team  *Team        `json:"team,omitempty"`
	Left  *DisplayNode `json:"left,omitempty"`
	Right *DisplayNode `json:"right,omitempty"`
}

func (bt *Bracket) ToDisplayNode(node *BracketNode) *DisplayNode {
	if node == nil {
		return nil
	}
	return &DisplayNode{
		Team:  node.Team,
		Left:  bt.ToDisplayNode(node.Left),
		Right: bt.ToDisplayNode(node.Right),
	}
}
