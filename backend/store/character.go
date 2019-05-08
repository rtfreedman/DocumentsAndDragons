package store

import (
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

func getCharacterFromCollection(ctx context.Context, characterIdentifier string, collection *mongo.Collection) (c *Character, err error) {
	// allocate memory
	c = &Character{}
	// get that objectid
	id, err := primitive.ObjectIDFromHex(characterIdentifier)
	if err != nil {
		return
	}
	// find the character in the collection
	filter := bson.D{{"_id", id}}
	result := collection.FindOne(ctx, filter)
	// decode into the character, if there's an error give it back
	err = result.Decode(c)
	return
}

// FindCharacter finds a character and returns that character object
func FindCharacter(ctx context.Context, characterIdentifier string) (c *Character, err error) {
	// if we have an already built character return that object from mongo
	if c, err = getCharacterFromCollection(ctx, characterIdentifier, characterCollection); err != nil {
		// if we get an error we're going to try to rebuild that character from the base character
		c, err = RebuildCharacter(ctx, characterIdentifier)
	}
	return
}

// RebuildCharacter will build a character from the base character object again and return the result
func RebuildCharacter(ctx context.Context, characterIdentifier string) (c *Character, err error) {
	// get the character from the base character collection so we can start rebuilding
	if c, err = getCharacterFromCollection(ctx, characterIdentifier, baseCharacterCollection); err != nil {
		// that shit don't exist
		return
	}
	// TODO: ensure we don't address multiple ID/collection pairs within the same block (so we don't end up in an infinite loop)
	seenIDs := map[primitive.ObjectID]string{}
	_ = seenIDs
	// we can build him, we have the technology
	// TODO: ensure the ordering here is right
	// TODO: Retrieve race changes
	// TODO: Retrieve background changes
	// TODO: Retrieve class changes
	// TODO: Retrieve spell changes
	// TODO: Retrieve ability changes
	// TODO: Retrieve item changes
	for index, item := range c.Inventory.Items {
		// retrieve information about the item
		result := itemCollection.FindOne(ctx, bson.D{{"_id", item.ID}})
		var tmpItem Item
		result.Decode(&tmpItem)
		// add the information to the character object that isn't custom populated
		_ = index
		// c.Inventory.Items[index].Union(tmpItem) // TODO: implement union
		if item.Equipped {
			// run equip
		}
	}
	return
}
