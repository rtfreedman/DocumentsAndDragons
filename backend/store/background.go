package store

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// AddBackground adds a background to a character
func (c *Character) AddBackground(b Background) (err error) {
	if err = c.update(b.Update); err != nil {
		return
	}
	if err = c.aggregateUpdate(b.Aggregate); err != nil {
		return
	}
	return
}

// AddBackground adds a new background to mongo
func AddBackground(doc *Background) (err error) {
	var ok bool
	if result, err := backgroundCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}
