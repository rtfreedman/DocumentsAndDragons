package store

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/mongo"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

func getCharacterFromCollection(characterIdentifier string, collection *mongo.Collection) (c *Character, err error) {
	// allocate memory
	c = &Character{}
	// get that objectid
	id, err := primitive.ObjectIDFromHex(characterIdentifier)
	if err != nil {
		return
	}
	// find the character in the collection
	filter := bson.D{{"_id", id}}
	result := collection.FindOne(backgroundContext, filter)
	// decode into the character, if there's an error give it back
	err = result.Decode(c)
	return
}

// FindCharacter finds a character and returns that character object
func FindCharacter(characterIdentifier string) (c *Character, err error) {
	// if we have an already built character return that object from mongo
	if c, err = getCharacterFromCollection(characterIdentifier, characterCollection); err != nil {
		// if we get an error we're going to try to rebuild that character from the base character
		c, err = RebuildCharacter(characterIdentifier)
	}
	return
}

// RebuildCharacter will build a character from the base character object again and return the result
func RebuildCharacter(characterIdentifier string) (c *Character, err error) {
	// get the character from the base character collection so we can start rebuilding
	// we get the character interface here because we want to layer the base character information
	// over the default info
	m := map[string]interface{}{}
	// get that objectid
	id, err := primitive.ObjectIDFromHex(characterIdentifier)
	if err != nil {
		return
	}
	// find the character in the collection
	filter := bson.D{{"_id", id}}
	result := baseCharacterCollection.FindOne(backgroundContext, filter)
	// decode into the character, if there's an error give it back
	if err = result.Decode(&m); err != nil {
		return
	}
	c = &Character{ID: id}
	if err = c.copyBaseCharacteristics(m); err != nil {
		return
	}
	// TODO: ensure we don't address multiple ID/collection pairs within the same block (so we don't end up in an infinite loop)
	seenIDs := map[primitive.ObjectID]string{}
	_ = seenIDs
	// we can build him, we have the technology
	// TODO: ensure the ordering here is right
	// TODO: Retrieve race changes
	// TODO: Retrieve background changes
	// TODO: Retrieve class changes
	// TODO: Retrieve spell changes
	// TODO: Retrieve ability changes
	// TODO: Retrieve item changes
	err = c.getItemsFromMap(m)
	// handle the inventory
	return
}

func (c *Character) getItemsFromMap(m map[string]interface{}) (err error) {
	// extract the items from the character's inventory from base characters
	val, ok := m["items"]
	if !ok {
		// if the inventory doesn't exist just return. it's not a big deal
		return
	}
	// comes back as a bson.A
	inventory, ok := val.(bson.A)
	if !ok {
		// we return an error here because we know the items exist, but the inventory is malformed
		return errors.New("inventory malformed")
	}
	// we iterate over the items int he inventory
	for _, itemInterface := range inventory {
		// and convert them eventually to a map that we can mold into the character inventory properly
		item, ok := itemInterface.(bson.D)
		if !ok {
			// TODO: log an error here
			continue
		}
		// now we add that item to the map
		if err = c.addItemToInventory(item); err != nil {
			// TODO: log an error here
			continue
		}
	}

	return errors.New("implementing")
}

func (c *Character) addItemToInventory(item bson.D) (err error) {
	itemMap := item.Map()
	var id primitive.ObjectID
	if idInterface, ok := itemMap["_id"]; !ok {
		return errors.New("no id in item")
	} else if id, ok = idInterface.(primitive.ObjectID); !ok {
		return errors.New("malformed id in item")
	}
	baseItem, err := FindItem(id)
	if err != nil {
		return
	}
	_ = baseItem
	return
}

func (c *Character) copyBaseCharacteristics(m map[string]interface{}) (err error) {
	// TODO: refactor all this crap if I figure out a programmatically clean way to do so
	var val interface{}
	var ok bool
	// mandatory elements
	if val, ok = m["user"]; !ok {
		return errors.New("no user associated with character")
	} else if c.User, ok = val.(primitive.ObjectID); !ok {
		return errors.New("user element not Object id on decode")
	}
	if val, ok = m["campaign"]; !ok {
		return errors.New("no campaign associated with character")
	} else if c.User, ok = val.(primitive.ObjectID); !ok {
		return errors.New("campaign element not Object id on decode")
	}
	// populated if not available
	if val, ok = m["name"]; !ok {
		c.Name = "Unnamed Character"
	} else if c.Name, ok = val.(string); !ok {
		c.Name = "Unnamed Character"
	}
	// optional []string
	if val, ok = m["advantages"]; ok {
		// we don't care if this is empty becase it doesn't exist regardless right now
		c.Advantages, _ = val.([]string)
	}
	if val, ok = m["disadvantages"]; ok {
		c.Disadvantages, _ = val.([]string)
	}
	if val, ok = m["proficient"]; ok {
		c.Proficient, _ = val.([]string)
	}
	// optional string
	if val, ok = m["deity"]; ok {
		c.Deity, _ = val.(string)
	}
	if val, ok = m["patron"]; ok {
		c.Patron, _ = val.(string)
	}
	if val, ok = m["gender"]; ok {
		c.Gender, _ = val.(string)
	}
	if val, ok = m["eyeColor"]; ok {
		c.EyeColor, _ = val.(string)
	}
	if val, ok = m["hair"]; ok {
		c.Hair, _ = val.(string)
	}
	if val, ok = m["alignment"]; ok {
		c.Alignment, _ = val.(string)
	}
	// optional ints
	if val, ok = m["armorClass"]; ok {
		c.ArmorClass, _ = val.(int)
	}
	if val, ok = m["baseHitpoints"]; ok {
		c.BaseHitpoints, _ = val.(int)
	}
	if val, ok = m["height"]; ok {
		c.Height, _ = val.(int)
	}
	if val, ok = m["weight"]; ok {
		c.Weight, _ = val.(int)
	}
	if val, ok = m["age"]; ok {
		c.Age, _ = val.(int)
	}
	if val, ok = m["xp"]; ok {
		c.XP, _ = val.(int)
	}
	// ability scores
	if val, ok = m["str"]; ok {
		if c.STR, ok = val.(int); !ok {
			c.STR = 10
		}
	} else {
		c.STR = 10
	}
	if val, ok = m["dex"]; ok {
		if c.DEX, ok = val.(int); !ok {
			c.DEX = 10
		}
	} else {
		c.DEX = 10
	}
	if val, ok = m["con"]; ok {
		if c.CON, ok = val.(int); !ok {
			c.CON = 10
		}
	} else {
		c.CON = 10
	}
	if val, ok = m["int"]; ok {
		if c.INT, ok = val.(int); !ok {
			c.INT = 10
		}
	} else {
		c.INT = 10
	}
	if val, ok = m["wis"]; ok {
		if c.WIS, ok = val.(int); !ok {
			c.WIS = 10
		}
	} else {
		c.WIS = 10
	}
	if val, ok = m["cha"]; ok {
		if c.CHA, ok = val.(int); !ok {
			c.CHA = 10
		}
	} else {
		c.CHA = 10
	}
	return
}
