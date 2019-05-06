package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var backgroundCollection *mongo.Collection

// AddBackground adds a background to mongo
func AddBackground(ctx context.Context, background *Background) (err error) {
	result, err := backgroundCollection.InsertOne(ctx, background)
	if err != nil {
		return
	}
	var ok bool
	if background.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		err = errors.New("primitive.ObjectID not returned by InsertOne")
	}
	return
}
