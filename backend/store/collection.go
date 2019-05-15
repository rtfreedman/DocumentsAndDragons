package store

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/rtfreedman/DocumentsAndDragons/backend/util"
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

var characterCollections = map[string]chan bool{}

func initCollections() {
	abilityCollection = db.Collection("abilities")
	backgroundCollection = db.Collection("backgrounds")
	classCollection = db.Collection("classes")
	baseCharacterCollection = db.Collection("baseCharacters")
	characterCollection = db.Collection("characters")
	itemCollection = db.Collection("items")
	raceCollection = db.Collection("races")
	spellCollection = db.Collection("spells")
	for i := 0; i < 100; i++ {
		s, err := util.RandomHex(20)
		// because mongo collections can't begin with a number
		s = "c" + s
		if err != nil {
			panic(err.Error())
		}
		characterCollections[s] = make(chan bool, 1)
	}
}
