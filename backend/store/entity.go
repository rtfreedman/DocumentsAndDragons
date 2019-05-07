package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// AddEntity will add an entity to the colleciton provided
func AddEntity(ctx context.Context, doc *Document, collectionString string) (err error) {
	collection, ok := collectionMap[collectionString]
	if !ok {
		return errors.New("Bad collection name: " + collectionString)
	}
	if result, err := collection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}
