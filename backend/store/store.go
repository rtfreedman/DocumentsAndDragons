package store

import (
	"context"
	"fmt"
	"log"
	"syscall"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/crypto/ssh/terminal"
)

var db *mongo.Database
var abilityCollection *mongo.Collection
var backgroundCollection *mongo.Collection
var classCollection *mongo.Collection
var characterCollection *mongo.Collection
var baseCharacterCollection *mongo.Collection
var itemCollection *mongo.Collection
var raceCollection *mongo.Collection
var spellCollection *mongo.Collection
var campaignCollection *mongo.Collection

var collectionMap map[string]*mongo.Collection

var ctx = context.Background()

func init() {
	fmt.Println("Initiating mongo driver...")
	var err error
	db, err = configDB(ctx, false, "", "127.0.0.1", "27041", "dnddb")
	if err != nil {
		log.Fatal(err.Error())
	}
	abilityCollection = db.Collection("abilities")
	backgroundCollection = db.Collection("backgrounds")
	classCollection = db.Collection("classes")
	characterCollection = db.Collection("characters")
	baseCharacterCollection = db.Collection("baseCharacters")
	itemCollection = db.Collection("items")
	raceCollection = db.Collection("races")
	spellCollection = db.Collection("spells")
	// TODO: remove when no longer deving
	// ~~~~~NO SERIOUSLY REMOVE THIS BITCH~~~~~
	defer itemCollection.DeleteMany(ctx, bson.D{{}})
	defer characterCollection.DeleteMany(ctx, bson.D{{}})
	defer baseCharacterCollection.DeleteMany(ctx, bson.D{{}})
	heavyArmor := Item{
		Name: "Heavy Armor",
		Equip: bson.A{
			bson.D{
				{"$set", bson.D{{"armorClass", 15}}},
			},
		},
	}
	err = AddItem(&heavyArmor)
	if err != nil {
		panic(err.Error())
	}
	onehex, err := primitive.ObjectIDFromHex("111111111111111111111111")
	if err != nil {
		panic(err.Error())
	}
	character := Character{
		User:       onehex,
		Campaign:   onehex,
		Name:       "Rorik Ironforge",
		ArmorClass: 10,
		STR:        24,
		CON:        24,
		DEX:        10,
		CHA:        14,
		WIS:        14,
		Items: []Item{
			Item{
				ID:       heavyArmor.ID,
				Equipped: true,
			},
			Item{
				ID:    heavyArmor.ID,
				Count: 2,
			},
		},
	}
	err = AddCharacter(&character)
	if err != nil {
		panic(err.Error())
	}
	if c, err := FindCharacter(character.ID.Hex()); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(c.ArmorClass)
	}
	panic(":)")
}

func getPassword() string {
	fmt.Print("Mongo Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatal("Error getting mongo terminal password")
	}
	password := string(bytePassword)
	return password
}

func configDB(ctx context.Context, hasAuth bool, userName string, host string, port string, dbName string) (db *mongo.Database, err error) {
	var uri string
	if hasAuth {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", userName, getPassword(), host, port)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s", host, port)
	}
	client, err := mongo.NewClient(uri)
	if err != nil {
		return
	}
	err = client.Connect(ctx)
	db = client.Database(dbName)
	return
}
