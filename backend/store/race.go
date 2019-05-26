package store

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// AddRace adds a race to a character
func (c *Character) AddRace(r Race) (err error) {
	if err = c.update(r.Update); err != nil {
		return
	}
	if err = c.aggregateUpdate(r.Aggregate); err != nil {
		return
	}
	return
}

// AddRace adds a new race to mongo
func AddRace(doc *Race) (err error) {
	var ok bool
	if result, err := raceCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}
