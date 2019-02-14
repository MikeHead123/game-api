package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	."./game"

)
var controller = &Controller{Repository: Repository{}};

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/games", controller.GetGames).Methods("GET")
	router.HandleFunc("/games", controller.AddGame).Methods("POST")
	router.HandleFunc("/games", controller.UpdateGame).Methods("PUT")
	router.HandleFunc("/games", controller.DeleteGame).Methods("DELETE")
	// r.HandleFunc("/movies/{id}", FindGame).Methods("GET")
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
