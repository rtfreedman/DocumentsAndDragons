package store

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// AddSpell adds a new spell to mongo
func AddSpell(doc *Spell) (err error) {
	var ok bool
	if result, err := spellCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}

// UseSpell will use a spell
func (c *Character) UseSpell(s *Spell) (err error) {
	if err = c.update(s.Cast); err != nil {
		return
	}
	if err = c.aggregateUpdate(s.CastAggregate); err != nil {
		return
	}
	return
}
