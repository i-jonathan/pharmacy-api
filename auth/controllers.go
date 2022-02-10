package auth

import (
	"Pharmacy/account"
	"Pharmacy/core"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
)

var (
	client = core.GetRedisDB()
	store  *redisstore.RedisStore
	db     = core.GetDB()
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	var err error
	store, err = redisstore.NewRedisStore(context.Background(), client)

	if err != nil {
		log.Fatal("Failed at creating a redis store: ", err)
	}

	store.KeyPrefix("session_")

	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 1080 * 60,
	})
}

//Login
//@Summary      Login
//@Description  Log users in
//@Tags         auth
//@Accept       json
//@Param        login  body  credentials  true  "login"
//@Produce      json
//@Success      200
//@Failure      400  {object}  core.ErrorResponse
//@Failure      401  {object}  core.ErrorResponse
//@Router       /auth/login [post]
func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "chocolate_chip")
	// authenticate user

	var cred credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Check message body"})
		return
	}

	if cred.Email == "" || cred.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Email or Password can't be empty"})
		return
	}

	var user account.User
	db.Find(&user, "email = ?", strings.ToLower(cred.Email))

	// Verify tat password is correct
	expectedPass := user.Password
	correct, err := account.ComparePassword(cred.Password, expectedPass)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !correct {
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(core.ErrorResponse{Message: "Check Entered details."})
		return
	}
	// Save session and set user as authenticated
	session.Values["authenticated"] = true
	session.Values["email"] = user.Email
	session.Save(r, w)
}

//Logout
//@Summary      Logout
//@Description  Log users out
//@Tags         auth
//@Success      200
//@Router       /auth/logout [post]
func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "chocolate_chip")

	session.Values["authenticated"] = false
	session.Save(r, w)
}
