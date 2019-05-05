package store

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// FindCharacter finds a character and returns that character object
func FindCharacter(characterIdentifier string) (c *Character, err error) {
	c = &Character{}
	id, err := primitive.ObjectIDFromHex(characterIdentifier)
	if err != nil {
		return
	}
	filter := bson.D{{"_id", id}}
	result := characters.FindOne(context.Background(), filter)
	err = result.Decode(c)
	return
}

// UseAbility uses a character's ability
func (c *Character) UseAbility(ability string) (err error) {
	// TODO: validate user/dm
	// characters.UpdateOne(context.Background(), bson.D{{"_id": c.ID}})
	return
}
