package inventory

import "github.com/gorilla/mux"

func Router(router *mux.Router) *mux.Router {
	router.HandleFunc("/all", getAllItems).Methods("GET")
	router.HandleFunc("/{id}", getItem).Methods("GET")
	router.HandleFunc("/add", addItem).Methods("POST")
	router.HandleFunc("/sell-item", sellItem).Methods("POST")
	return router
}
