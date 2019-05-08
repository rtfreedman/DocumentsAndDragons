package store

import (
	"context"
	"errors"

	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// AddAbility adds a new ability to mongo
func AddAbility(ctx context.Context, doc *Ability) (err error) {
	var ok bool
	if result, err := abilityCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}

// AddBackground adds a new background to mongo
func AddBackground(ctx context.Context, doc *Background) (err error) {
	var ok bool
	if result, err := backgroundCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}

// AddClass adds a new class to mongo
func AddClass(ctx context.Context, doc *Class) (err error) {
	var ok bool
	if result, err := classCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}

// AddCharacter adds a new character to mongo
func AddCharacter(ctx context.Context, doc *Character) (err error) {
	var ok bool
	if result, err := baseCharacterCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}

// AddItem adds a new item to mongo
func AddItem(ctx context.Context, doc *Item) (err error) {
	var ok bool
	if result, err := itemCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}

// AddRace adds a new race to mongo
func AddRace(ctx context.Context, doc *Race) (err error) {
	var ok bool
	if result, err := raceCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}

// AddSpell adds a new spell to mongo
func AddSpell(ctx context.Context, doc *Spell) (err error) {
	var ok bool
	if result, err := spellCollection.InsertOne(ctx, doc); err != nil {
		return err
	} else if doc.ID, ok = result.InsertedID.(primitive.ObjectID); !ok {
		return errors.New("bad insert id returned")
	}
	return
}
