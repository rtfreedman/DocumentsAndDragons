package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var characterCollection *mongo.Collection
var baseCharacterCollection *mongo.Collection

// FindCharacter finds a character and returns that character object
func FindCharacter(characterIdentifier string) (c *Character, err error) {
	c = &Character{}
	id, err := primitive.ObjectIDFromHex(characterIdentifier)
	if err != nil {
		return
	}
	filter := bson.D{{"_id", id}}
	result := characterCollection.FindOne(context.Background(), filter)
	err = result.Decode(c)
	return
}

// UseAbility uses a character's ability
func (c *Character) UseAbility(ability string) (err error) {
	// TODO: validate user/dm
	// characters.UpdateOne(context.Background(), bson.D{{"_id": c.ID}})
	return
}

// EquipItem equips an item to a character
func (c *Character) EquipItem(ctx context.Context, item Item) {
	for _, d := range item.Equip {
		characterCollection.UpdateOne(ctx, bson.D{{"_id", c.ID}}, d)
	}
}

// AddCharacter adds a character to mongo
func AddCharacter(ctx context.Context, character *Character) (err error) {
	result, err := characterCollection.InsertOne(ctx, character)
	if err != nil {
		return
	}
	var ok bool
	if character.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		err = errors.New("primitive.ObjectID not returned by InsertOne")
	}
	return
}
