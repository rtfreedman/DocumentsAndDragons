package store

import (
	"errors"
	"fmt"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// FindItem will find an item according to the objectid supplied
func FindItem(id primitive.ObjectID) (item *Item, err error) {
	item = new(Item)
	result := itemCollection.FindOne(ctx, bson.D{{"_id", id}})
	err = result.Decode(&item)
	return
}

// FindItemFromString is a wrapper around FindItem that takes a string instead of an objectid
func FindItemFromString(idString string) (item *Item, err error) {
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(idString); err != nil {
		return nil, err
	}
	return FindItem(id)
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

// EquipItem will equip a character's item
func (c *Character) EquipItem(inventoryID int) (err error) {
	for _, item := range c.Items {
		if item.InventoryID == inventoryID {
			return c.equipItem(item)
		}
	}
	return errors.New("not implemented")
}

func (c *Character) equipItem(item Item) (err error) {
	for _, equip := range item.Equip {
		equipD, ok := equip.(bson.D)
		if !ok {
			return errors.New("bad equip for item")
		}
		fmt.Println(equipD)
		fmt.Println(c.ID)
		// TODO: aggregate don't update
		// if result, err := characterCollection.UpdateOne(ctx, bson.D{{"_id", c.ID}}, equipD); err != nil {
		// 	return err
		// } else if result.MatchedCount != 1 {
		// 	return errors.New("no item found")
		// }
	}
	return
}

// AddItemToInventory adds an item to a character's inventory. Expects the identifier associated with the item
func (c *Character) AddItemToInventory(itemID string) (err error) {
	var id primitive.ObjectID
	if id, err = primitive.ObjectIDFromHex(itemID); err != nil {
		return
	}
	return c.addItemToInventory(bson.D{{"_id", id}})
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
	// for each customization in the itemMap we overwrite the value in the base Item
	for key, value := range itemMap {
		// no _id case here because that is what we used to query on so it's present
		switch key {
		case "inventoryID":
			if val, ok := value.(int); ok {
				baseItem.InventoryID = val
			}
		case "name":
			if val, ok := value.(string); ok {
				baseItem.Name = val
			}
		case "charges":
			if val, ok := value.(int); ok {
				baseItem.Charges = val
			}
		case "maxCharges":
			if val, ok := value.(int); ok {
				baseItem.MaxCharges = val
			}
		case "description":
			if val, ok := value.(string); ok {
				baseItem.Description = val
			}
		case "rechargeCondition":
			if val, ok := value.(int); ok {
				baseItem.RechargeCondition = val
			}
		case "attunementRequired":
			if val, ok := value.(bool); ok {
				baseItem.AttunementRequired = val
			}
		case "attuned":
			if val, ok := value.(bool); ok {
				baseItem.Attuned = val
			}
		case "price":
			if val, ok := value.(int); ok {
				baseItem.Price = val
			}
		case "count":
			if val, ok := value.(int); ok {
				baseItem.Count = val
			}
		case "weight":
			if val, ok := value.(int); ok {
				baseItem.Weight = val
			}
		case "equipped":
			if val, ok := value.(bool); ok {
				baseItem.Equipped = val
			}
		case "equipPriority":
			if val, ok := value.(int); ok {
				baseItem.EquipPriority = val
			}
		case "stackable":
			if val, ok := value.(bool); ok {
				baseItem.Stackable = val
			}
		case "equip":
			if val, ok := value.(bson.A); ok {
				baseItem.Equip = val
			}
		case "use":
			if val, ok := value.(bson.A); ok {
				baseItem.Use = val
			}
		}
	}
	c.Items = append(c.Items, *baseItem)
	return
}
