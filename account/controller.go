package account

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getUser(r *http.Request, w http.ResponseWriter) {
	var user User
	db := InitDatabase()
	params := mux.Vars(r)
	username := params["username"]

	db.Find(&user, "username = ?", username)
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
	}
	// TODO log access
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	return
}

func getAllUsers(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db := InitDatabase()
	db.Find(&users)

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

func postUser(r *http.Request, w http.ResponseWriter) {
	// TODO check user details
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(user.Password) < 8 {
		return
	}

	// TODO hash password
	// TODO save user
}
