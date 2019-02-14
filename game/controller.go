package game

import (
	"encoding/json"
	"log"
	"net/http"
)

type Controller struct {
	Repository Repository
}

func (c *Controller) GetGames(w http.ResponseWriter, r *http.Request) {
	games := c.Repository.GetGames()
	log.Println(games)
	data, _ := json.Marshal(games)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

func (c *Controller) AddGame(w http.ResponseWriter, r *http.Request) {

	var game Game
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&game)
    if err != nil {
		log.Fatalln("Error Addgame data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
    }
	success := c.Repository.Addgame(game)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

func (c *Controller) UpdateGame(w http.ResponseWriter, r *http.Request) {
	var game Game
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&game)
    if err != nil {
		log.Fatalln("Error UpdateGame data", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
    }

	success := c.Repository.UpdateGame(game)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return
}

func (c *Controller) DeleteGame(w http.ResponseWriter, r *http.Request) {

	var game Game
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&game)

	success := c.Repository.DeleteGame(game)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	return
}
