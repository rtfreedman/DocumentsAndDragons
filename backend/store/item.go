package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var itemCollection *mongo.Collection

// AddItem adds an item to mongo
func AddItem(ctx context.Context, item *Item) (err error) {
	result, err := itemCollection.InsertOne(ctx, item)
	var ok bool
	if item.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		err = errors.New("primitive.ObjectID not returned by InsertOne")
	}
	return
}
