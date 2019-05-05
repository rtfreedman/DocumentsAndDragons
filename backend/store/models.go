package store

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Ability represents one of a character's abilities
type Ability struct {
	ID                primitive.ObjectID `bson:"_id"`
	Name              string             `bson:"name"`
	Description       string             `bson:"description"`
	Charges           int                `bson:"charges"`
	MaxCharges        int                `bson:"maxCharges"`
	RechargeCondition int                `bson:"rechargeCondition"`
	Active            bool               `bson:"active"` // vs passive
	Use               Filter             `bson:"use"`
}

// Background represents a character's background
type Background struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Alter       Filter             `bson:"alter"`
}

// Character is what is returned back to the user. This object is derived from the Base Character and all other filters applied
type Character struct {
	ID              primitive.ObjectID `bson:"_id"`
	User            primitive.ObjectID `bson:"user"`
	Campaign        primitive.ObjectID `bson:"campaign"`
	Name            string             `bson:"name"`
	AbilityScores   map[string]int     `bson:"abilityScores"`
	Proficient      map[string]bool    `bson:"proficient"` // proficiencies is derived from proficient
	BaseHitpoints   int                `bson:"baseHitpoints"`
	Deity           string             `bson:"deity"`
	Patron          string             `bson:"patron"`
	Height          int                `bson:"height"`
	Weight          int                `bson:"weight"`
	Age             int                `bson:"age"`
	Gender          string             `bson:"gender"`
	EyeColor        string             `bson:"eyeColor"`
	Hair            string             `bson:"hair"`
	Alignment       string             `bson:"alignment"`
	XP              int                `bson:"xp"`
	Race            Race               `bson:"race"`
	Background      Background         `bson:"background"`
	Classes         map[string]Class   `bson:"class"`
	Inventory       []Item             `bson:"inventory"`
	Abilities       map[string]Ability `bson:"abilities"`
	Spells          map[string]Spell   `bson:"spells"`
	SpellSlots      map[int]int        `bson:"spellSlots"`
	AvailableSpells map[int]int        `bson:"availableSpells"`
}

// Class represents a character's class
type Class struct {
	ID    primitive.ObjectID `bson:"_id"`
	Level int                `bson:"level"`
	Name  string             `bson:"name"`
	Alter map[int]Filter     `bson:"alter"` // changes per level
}

// Filter represents the changes that take place on the character object when an object is used or equipped
type Filter struct {
	Filter bson.D
	Update bson.D
}

// Item represents one of a character's items
type Item struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Name               string             `bson:"name"`
	Charges            int                `bson:"charges"`
	MaxCharges         int                `bson:"maxCharges"`
	Description        string             `bson:"description"`
	RechargeCondition  int                `bson:"rechargeCondition"`
	AttunementRequired bool               `bson:"attunementRequired"`
	Attuned            bool               `bson:"attuned"`
	Price              int                `bson:"price"`
	Count              int                `bson:"count"`
	Weight             int                `bson:"weight"`
	Use                Filter             `bson:"use"`
	Equip              Filter             `bson:"equip"`
	Stackable          bool               `bson:"stackable"`
}

// Race represents a character's race
type Race struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Alter       Filter             `bson:"alter"`
}

// Spell represents an individual spell that a character has the ability to learn
type Spell struct {
	ID            primitive.ObjectID   `bson:"_id"`
	Name          string               `bson:"name"`
	Description   string               `bson:"description"`
	CastingTime   string               `bson:"castingTime"`
	Classes       []primitive.ObjectID `bson:"classes"`
	Components    string               `bson:"components"`
	Concentration bool                 `bson:"concentration"`
	Duration      string               `bson:"duration"`
	Level         int                  `bson:"level"`
	Use           Filter               `bson:"use"`
}
