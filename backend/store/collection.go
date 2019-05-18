package store

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

var abilityCollection *mongo.Collection
var backgroundCollection *mongo.Collection
var classCollection *mongo.Collection
var baseCharacterCollection *mongo.Collection
var itemCollection *mongo.Collection
var raceCollection *mongo.Collection
var spellCollection *mongo.Collection
var campaignCollection *mongo.Collection
var characterCollection *mongo.Collection

func initCollections() {
	abilityCollection = db.Collection("abilities")
	backgroundCollection = db.Collection("backgrounds")
	classCollection = db.Collection("classes")
	baseCharacterCollection = db.Collection("baseCharacters")
	characterCollection = db.Collection("characters")
	itemCollection = db.Collection("items")
	raceCollection = db.Collection("races")
	spellCollection = db.Collection("spells")
}
