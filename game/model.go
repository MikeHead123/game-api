package game

import "gopkg.in/mgo.v2/bson"

type Game struct {
	ID     bson.ObjectId `bson:"_id"`
	Title  string        `json:"title"`
	Description string        `json:"description"`
	Year   int32         `json:"year"`
}

type Games []Game
