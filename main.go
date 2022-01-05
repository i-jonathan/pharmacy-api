package main

import (
	"Pharmacy/account"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("starting server")
	router := mux.NewRouter()
	account.Router(router.PathPrefix("/account").SubRouter())

	err := http.ListenAndServe(":9560", router)
	if err != nil {
		log.Println(err)
	}
}
