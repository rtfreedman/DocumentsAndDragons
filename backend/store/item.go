package store

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// FindItem will find an item according to the objectid supplied
func FindItem(id primitive.ObjectID) (item *Item, err error) {
	item = new(Item)
	result := itemCollection.FindOne(backgroundContext, bson.D{{"_id", id}})
	err = result.Decode(&item)
	return
}

// FindItemFromString is a wrapper around FindItem that takes a string instead of an objectid
func FindItemFromString(idString string) (item *Item, err error) {
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(idString); err != nil {
		return nil, err
	}
	return FindItem(id)
}
