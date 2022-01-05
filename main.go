package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting server")
	router := mux.NewRouter()

	err := http.ListenAndServe(":9560", router)
	if err != nil {
		log.Println(err)
	}
}
