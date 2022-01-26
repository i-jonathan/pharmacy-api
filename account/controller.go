package account

import (
	"Pharmacy/core"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var db = core.GetDB()

//Get User
//@Summary      Get user
//@Description  Get user by Username
//@Tags         user
//@Produce      json
//@Param        username  path      string  true  "username"
//@Success      200  {object}  User
//@Failure      404       {object}  core.ErrorResponse
//@Router       /account/{username} [get]
func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
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
	return
}

//Get All Users
//@Summary      Get users
//@Description  Get all users
//@Tags         user
//@Produce      json
//@Success      200  {object}  core.Response{[]data=User}
//@Success      204  {object}  core.Response{[]data=User}
//@Router       /account/all [get]
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	var count int64
	db.Scopes(core.Paginate(r)).Find(&users)
	db.Model(&User{}).Count(&count)

	if len(users) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	page, prev, next := core.ResponseData(int(count), r)
	response := core.Response{
		Previous: prev,
		Next:     next,
		Page:     page,
		Count:    count,
		Data:     users,
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	return
}

//Set User
//@Summary      Set user
//@Description  Create new user account
//@Tags         user
//@Accept       json
//@Produce      json
//@Success      200       {object}  User
//@Failure      400  {object}  core.ErrorResponse
//@Failure      500
//@Router       /account [post]
func postUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// TODO check user details
	var user User
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
