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
	Team   *Team
	Left   *BracketNode
	Right  *BracketNode
	Parent *BracketNode
}

type Bracket struct {
	Root *BracketNode
}

type StorableBracketNode struct {
	Team    string               `bson:"team,omitempty" json:"team,omitempty"`
	TeamID  primitive.ObjectID   `bson:"teamId,omitempty" json:"teamId,omitempty"`
	Seeding int                  `bson:"seeding,omitempty" json:"seeding,omitempty"`
	Left    *StorableBracketNode `bson:"left,omitempty" json:"left,omitempty"`
	Right   *StorableBracketNode `bson:"right,omitempty" json:"right,omitempty"`
}

type StorableBracket struct {
	Event string               `bson:"event,omitempty" json:"event,omitempty"`
	Root  *StorableBracketNode `bson:"root,omitempty" json:"root,omitempty"`
}

func GetBracket(eventSlug string) StorableBracket {
	var result StorableBracket
	err := configs.BracketsCollection.FindOne(context.TODO(), bson.D{{Key: "event", Value: eventSlug}}).Decode(&result)
	if err != nil {
		panic(err)
	}

	return result
}

func InsertBracket(newBracket StorableBracket) {
	_, err := configs.BracketsCollection.InsertOne(context.TODO(), newBracket)
	if err != nil {
		panic(err)
	}
}

// Convert Bracket to StorableBracket
func (b *Bracket) ToStorable(eventSlug string) *StorableBracket {
	return &StorableBracket{
		Root:  convertToStorableNode(b.Root),
		Event: eventSlug,
	}
}

func convertToStorableNode(node *BracketNode) *StorableBracketNode {
	if node == nil {
		return nil
	}

	storableNode := &StorableBracketNode{
		Left:  convertToStorableNode(node.Left),
		Right: convertToStorableNode(node.Right),
	}

	if node.Team != nil {
		storableNode.TeamID = node.Team.ID
		storableNode.Team = node.Team.Name
		storableNode.Seeding = node.Team.Seeding
	}

	return storableNode
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
	matchups := getBracket(teams)

	bracket := createBracketTree(matchups)
	// printBracketTree(bracket.Root, 0)
	StorableBracket := *bracket.ToStorable(eventSlug)
	InsertBracket(StorableBracket)
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

func createBracketTree(matchups [][]*Team) Bracket {
	if len(matchups) == 0 {
		return Bracket{Root: nil}
	}

	// Create leaf nodes
	leaves := make([]*BracketNode, 0, len(matchups)*2)
	for _, matchup := range matchups {
		leaves = append(leaves, &BracketNode{Team: matchup[0]})
		leaves = append(leaves, &BracketNode{Team: matchup[1]})
	}

	// Build the tree bottom-up
	for len(leaves) > 1 {
		parents := make([]*BracketNode, 0, len(leaves)/2)
		for i := 0; i < len(leaves); i += 2 {
			parent := &BracketNode{
				Left:  leaves[i],
				Right: leaves[i+1],
			}
			leaves[i].Parent = parent
			leaves[i+1].Parent = parent
			parents = append(parents, parent)
		}
		leaves = parents
	}

	return Bracket{Root: leaves[0]}
}

func PrintBracketTree(node *BracketNode, depth int) {
	if node == nil {
		return
	}

	PrintBracketTree(node.Right, depth+1)

	for i := 0; i < depth; i++ {
		fmt.Print("    ")
	}
	if node.Team == nil {
		fmt.Println("[-]")
	} else {
		fmt.Printf("[%s (Seed %d)]\n", node.Team.Name, node.Team.Seeding)
	}

	PrintBracketTree(node.Left, depth+1)
}
