package store

import (
	"context"
	"fmt"
	"log"
	"syscall"

	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/crypto/ssh/terminal"
)

var db *mongo.Database
var abilityCollection *mongo.Collection
var backgroundCollection *mongo.Collection
var characterCollection *mongo.Collection
var itemCollection *mongo.Collection
var raceCollection *mongo.Collection
var spellCollection *mongo.Collection

var collectionMap map[string]*mongo.Collection

func init() {
	fmt.Println("Initiating mongo driver...")
	var err error
	db, err = configDB(context.Background(), false, "", "127.0.0.1", "27041", "dnddb")
	if err != nil {
		log.Fatal(err.Error())
	}
	abilityCollection = db.Collection("ability")
	spellCollection = db.Collection("spell")
	raceCollection = db.Collection("race")
	backgroundCollection = db.Collection("background")
	itemCollection = db.Collection("items")
	characterCollection = db.Collection("characters")
	collectionMap = map[string]*mongo.Collection{
		"ability":    abilityCollection,
		"background": backgroundCollection,
		"character":  characterCollection,
		"item":       itemCollection,
		"race":       raceCollection,
		"spell":      spellCollection,
	}
	character := Document{
		Name:       "Rorik Ironforge",
		ArmorClass: 15,
		AbilityScores: AbilityScores{
			STR: 24,
			CON: 24,
			DEX: 10,
			CHA: 14,
			WIS: 14,
		},
	}
	fmt.Println(AddCharacter(context.Background(), &character))
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
