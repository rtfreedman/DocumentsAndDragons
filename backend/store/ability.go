package store

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// AddAbility adds an ability to a character
func (c *Character) AddAbility(a *Ability) (err error) {
	if err = c.update(a.Update); err != nil {
		return
	}
	if err = c.aggregateUpdate(a.Aggregate); err != nil {
		return
	}
	return
}

// AddAbility adds a new ability to mongo
func AddAbility(doc *Ability) (err error) {
	var ok bool
	if result, err := abilityCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}
