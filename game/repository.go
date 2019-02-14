package game

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Repository struct{}

const SERVER = "localhost:27017"

const DBNAME = "gameStore"

const DOCNAME = "games"

func (r Repository) GetGames() Games {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Games{}
	c.Find(nil).All(&results)

	return results
}

func (r Repository) Addgame(game Game) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	game.ID = bson.NewObjectId()
	session.DB(DBNAME).C(DOCNAME).Insert(game)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) UpdateGame(game Game) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	session.DB(DBNAME).C(DOCNAME).UpdateId(game.ID, game)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (r Repository) DeleteGame(game Game) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()

	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(game.ID); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
