package auth

import (
	"Pharmacy/core"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore"
)

var (
	client = core.GetRedisDB()
	store  *redisstore.RedisStore
)

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

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "chocolate_chip")
	// authenticate user

	type credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var cred credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	if cred.Email == "" || cred.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user account.User
	db.Find(&user, "email = ?", strins.ToLower(cred.Email))

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

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "chocolate_chip")

	session.Values["authenticated"] = false
	session.Save(r, w)
}
