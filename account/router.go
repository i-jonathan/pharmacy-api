package account

import "github.com/gorilla/mux"

func Router(router *mux.Router) *mux.Router {
	router.HandleFunc("/all", getAllUsers).Methods("GET")
	router.HandleFunc("", postUser).Methods("POST")
	router.HandleFunc("/{username}", getUser).Methods("GET")
	return router
}
