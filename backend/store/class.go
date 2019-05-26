package store

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// AddClass adds a class to a character
func (c *Character) AddClass(l Class) (err error) {
	if err = c.update(l.Update); err != nil {
		return
	}
	if err = c.aggregateUpdate(l.Aggregate); err != nil {
		return
	}
	return
}

// AddClass adds a new class to mongo
func AddClass(doc *Class) (err error) {
	var ok bool
	if result, err := classCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}
