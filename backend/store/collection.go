package store

import (
	"errors"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
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

var characterCollections = map[string]chan bool{}

func initCollections() {
	abilityCollection = db.Collection("abilities")
	backgroundCollection = db.Collection("backgrounds")
	classCollection = db.Collection("classes")
	baseCharacterCollection = db.Collection("baseCharacters")
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

func checkoutCollection() (collection *mongo.Collection, err error) {
	t := time.After(25 * time.Second)
	for {
		for k := range characterCollections {
			select {
			case <-t:
				return nil, errors.New("timeout on collection get")
			case characterCollections[k] <- true:
				collection = db.Collection(k)
				return
			// this is so we don't get an issue with unschedulable goroutines whenever this is called, locking our program
			case <-time.After(1 * time.Microsecond):
				continue
			}
		}
	}
}

func clearCollection(collection *mongo.Collection) (err error) {
	_, err = collection.DeleteMany(ctx, bson.D{{}})
	return
}

func returnCollection(collection *mongo.Collection) {
	<-characterCollections[collection.Name()]
}
