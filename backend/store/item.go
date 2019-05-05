package store

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// Item represents an item in a character's inventory
type Item struct {
	ID                primitive.ObjectID
	Name              string
	Charges           int
	MaxCharges        int
	Description       string
	RechargeCondition int
	Attuned           bool
	Price             int
	Use               Filter
	Equip             Filter
	Unequip           Filter
}
