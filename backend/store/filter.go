package store

import (
	"github.com/mongodb/mongo-go-driver/bson"
)

// Filter represents the changes that take place on the character object when an object is used or equipped
type Filter struct {
	Filter bson.D
	Update bson.A
}
