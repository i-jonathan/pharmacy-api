package main

import (
	"Pharmacy/account"
	"Pharmacy/inventory"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//@title          Pharmacy First Steps
//@version        1.0
//@contact.name   Jonathan Farinloye
//@contact.email  farinloyejonathan@gmail.com
//@description    An API server for a not yet operational pharmacy
func main() {
	log.Println("starting server")
	router := mux.NewRouter()

	// documentation
	fs := http.FileServer(http.Dir("./docs"))
	router.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", fs))

	account.Router(router.PathPrefix("/account").Subrouter())
	inventory.Router(router.PathPrefix("/inventory").Subrouter())

	err := http.ListenAndServe(":9560", router)
	if err != nil {
		log.Println(err)
	}
}
