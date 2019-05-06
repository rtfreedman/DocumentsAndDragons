package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var abilityCollection *mongo.Collection

// AddAbility adds a spell to mongo
func AddAbility(ctx context.Context, ability *Ability) (err error) {
	result, err := abilityCollection.InsertOne(ctx, ability)
	var ok bool
	if ability.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		err = errors.New("primitive.ObjectID not returned by InsertOne")
	}
	return
}
