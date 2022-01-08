package account

import (
	"Pharmacy/core"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Get User
//@Summary      Get user
//@Description  Get user by Username
//@Tags         user
//@Produce      json
//@Param        username  path      string  true  "username"
//@Success      200       {object}  User
//@Failure      404       {object}  core.ErrorResponse
//@Router       /account/{username} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	var user User
	db := InitDatabase()
	params := mux.Vars(r)
	username := params["username"]

	db.Find(&user, "username = ?", username)

	if user.Username == "" {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Not found"})
		if err != nil {
			log.Println(err)
		}
		return
	}

	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
	}

	// TODO log access
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	return
}

//Get All Users
//@Summary      Get users
//@Description  Get all users
//@Tags         user
//@Produce      json
//@Success      200       {array}  User
//@Success	204 {array} User
//@Router /account/all [get]
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db := InitDatabase()
	db.Find(&users)

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	return
}

//Add User
//@Summary      Add user
//@Description  Create new user account
//@Tags         user
//@Produce      json
//@Success      200       {object}  User
//@Failure 400 {object} core.ErrorResponse
//@Failure 500
//@Router /account [post]
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
		w.WriteHeader(http.StatusBadRequest)
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
