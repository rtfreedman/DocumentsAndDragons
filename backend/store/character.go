package store

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// Character represents a character in Mongo
type Character struct {
	ID            primitive.ObjectID
	Level         int
	Name          string
	AbilityScores map[string]int
	Proficiencies map[string]int
	BaseHitpoints int
	Race          string
	Subrace       string
	Diety         string
	DeityStatus   string
	Languages     []string
	Height        string
	Weight        int
	Age           int
	Gender        string
	EyeColor      string
	Hair          string
	Alignment     string
	Background    string
	XP            int
	Inventory     []Item
	Abilities     []Ability
}
