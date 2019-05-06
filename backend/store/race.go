package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var raceCollection *mongo.Collection

// AddRace adds a race to mongo
func AddRace(ctx context.Context, race *Race) (err error) {
	result, err := raceCollection.InsertOne(ctx, race)
	var ok bool
	if race.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		err = errors.New("primitive.ObjectID not returned by InsertOne")
	}
	return
}
