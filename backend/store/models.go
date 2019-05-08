package store

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Ability represents one of a character's abilities
type Ability struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name,omitempty"`
	Description       string             `bson:"description,omitempty"`
	Charges           int                `bson:"charges,omitempty"`
	MaxCharges        int                `bson:"maxCharges,omitempty"`
	RechargeCondition int                `bson:"rechargeCondition,omitempty"`
	Active            bool               `bson:"active,omitempty"` // vs passive
	Use               bson.A             `bson:"use,omitempty"`
}

// AbilityScores represents the ability scores of a character
// not really necessary but is cleaner than using a map
type AbilityScores struct {
	STR int `bson:"str,omitempty"`
	DEX int `bson:"dex,omitempty"`
	CON int `bson:"con,omitempty"`
	INT int `bson:"int,omitempty"`
	WIS int `bson:"wis,omitempty"`
	CHA int `bson:"cha,omitempty"`
}

// Background represents a character's background
type Background struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Update      bson.A             `bson:"update,omitempty"`
}

// Class represents a character's class
type Class struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Level  int                `bson:"level,omitempty"`
	Update bson.A             `bson:"update,omitempty"`
}

// Character represents the shape of a character we expect out of mongo
type Character struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	User            primitive.ObjectID `bson:"user,omitempty"`
	Campaign        primitive.ObjectID `bson:"campaign,omitempty"`
	Name            string             `bson:"name,omitempty"`
	Advantages      []string           `bson:"disadvantages,omitempty"`
	Disadvantages   []string           `bson:"advantages,omitempty"`
	AbilityScores   AbilityScores      `bson:"abilityScores,omitempty"`
	ArmorClass      int                `bson:"armorClass,omitempty"`
	Proficient      []string           `bson:"proficient,omitempty"` // proficiencies is derived from proficient
	BaseHitpoints   int                `bson:"baseHitpoints,omitempty"`
	Deity           string             `bson:"deity,omitempty"`
	Patron          string             `bson:"patron,omitempty"`
	Height          int                `bson:"height,omitempty"`
	Weight          int                `bson:"weight,omitempty"`
	Age             int                `bson:"age,omitempty"`
	Gender          string             `bson:"gender,omitempty"`
	EyeColor        string             `bson:"eyeColor,omitempty"`
	Hair            string             `bson:"hair,omitempty"`
	Alignment       string             `bson:"alignment,omitempty"`
	XP              int                `bson:"xp,omitempty"`
	Race            Race               `bson:"race,omitempty"`
	Background      Background         `bson:"background,omitempty"`
	Classes         []Class            `bson:"class,omitempty"`
	Inventory       Inventory          `bson:"inventory,omitempty"`
	Abilities       []Ability          `bson:"abilities,omitempty"`
	Spells          Spells             `bson:"spells,omitempty"`
	SpellSlots      Spells             `bson:"spellSlots,omitempty"`
	AvailableSpells AvailableSpells    `bson:"availableSpells,omitempty"`
}

// Inventory is a character's inventory
type Inventory struct {
	Items []Item `bson:"items,omitempty"`
}

// Item represents one of a character's items
type Item struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	InventoryID        int                `bson:"inventoryID,omitempty"`
	Name               string             `bson:"name,omitempty"`
	Charges            int                `bson:"charges,omitempty"`
	MaxCharges         int                `bson:"maxCharges,omitempty"`
	Description        string             `bson:"description,omitempty"`
	RechargeCondition  int                `bson:"rechargeCondition,omitempty"`
	AttunementRequired bool               `bson:"attunementRequired,omitempty"`
	Attuned            bool               `bson:"attuned,omitempty"`
	Price              int                `bson:"price,omitempty"`
	Count              int                `bson:"count,omitempty"`
	Weight             int                `bson:"weight,omitempty"`
	Equipped           bool               `bson:"equipped,omitempty"`
	Stackable          bool               `bson:"stackable,omitempty"`
	Equip              bson.A             `bson:"equip,omitempty"`
	Use                bson.A             `bson:"use,omitempty"`
}

// Race represents a character's race
type Race struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Update      bson.A             `bson:"update,omitempty"`
}

// Spell represents an individual spell that a character has the ability to learn
type Spell struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	Description   string             `bson:"description,omitempty"`
	CastingTime   string             `bson:"castingTime,omitempty"`
	Components    string             `bson:"components,omitempty"`
	Concentration bool               `bson:"concentration,omitempty"`
	Duration      string             `bson:"duration,omitempty"`
	Level         int                `bson:"level,omitempty"`
	Cast          bson.A             `bson:"cast,omitempty"`
}

// Spells represent a character's spell shapes
type Spells struct {
	First   int `bson:"1,omitempty"`
	Second  int `bson:"2,omitempty"`
	Third   int `bson:"3,omitempty"`
	Fourth  int `bson:"4,omitempty"`
	Fifth   int `bson:"5,omitempty"`
	Sixth   int `bson:"6,omitempty"`
	Seventh int `bson:"7,omitempty"`
	Eighth  int `bson:"8,omitempty"`
	Ninth   int `bson:"9,omitempty"`
}

// AvailableSpells are the spells available to the character
type AvailableSpells struct {
	First   []primitive.ObjectID `bson:"1,omitempty"`
	Second  []primitive.ObjectID `bson:"2,omitempty"`
	Third   []primitive.ObjectID `bson:"3,omitempty"`
	Fourth  []primitive.ObjectID `bson:"4,omitempty"`
	Fifth   []primitive.ObjectID `bson:"5,omitempty"`
	Sixth   []primitive.ObjectID `bson:"6,omitempty"`
	Seventh []primitive.ObjectID `bson:"7,omitempty"`
	Eighth  []primitive.ObjectID `bson:"8,omitempty"`
	Ninth   []primitive.ObjectID `bson:"9,omitempty"`
}
