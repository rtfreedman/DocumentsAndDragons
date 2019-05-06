package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var spellCollection *mongo.Collection

// AddSpell adds a spell to mongo
func AddSpell(ctx context.Context, spell *Spell) (err error) {
	result, err := spellCollection.InsertOne(ctx, spell)
	var ok bool
	if spell.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		err = errors.New("primitive.ObjectID not returned by InsertOne")
	}
	return
}
