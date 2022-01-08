package account

import (
	"Pharmacy/core"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func getUser(w http.ResponseWriter, r *http.Request) {
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

func getAllUsers(w http.ResponseWriter, r *http.Request) {
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

func postUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	// TODO check user details
	var user User
	db := InitDatabase()
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(user.Password) < 8 {
		message := core.ErrorResponse{Message: "Password length"}
		_ = json.NewEncoder(w).Encode(message)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	// hash password
	user.Password, err = generatePasswordHash(user.Password)
	if err != nil {
		log.Println(err)
	}
	user.Username = strings.ToLower(strings.ReplaceAll(user.FullName, " ", ""))
	// save user
	db.Create(&user)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
	}
	return
}
