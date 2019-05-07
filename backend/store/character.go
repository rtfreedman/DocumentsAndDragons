package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// FindCharacter finds a character and returns that character object
func FindCharacter(characterIdentifier string) (c *Document, err error) {
	c = &Document{}
	id, err := primitive.ObjectIDFromHex(characterIdentifier)
	if err != nil {
		return
	}
	filter := bson.D{{"_id", id}}
	result := characterCollection.FindOne(context.Background(), filter)
	err = result.Decode(c)
	return
}

// AddCharacter adds a character to mongo
func AddCharacter(ctx context.Context, character *Document) (err error) {
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
