package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/interface/mux/helper"
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/service"
)

type accountController struct {
	svc service.AccountUseCase
}

func NewAccountController(s service.AccountUseCase) *accountController {
	return &accountController{s}
}

func (controller *accountController) FetchAccounts(w http.ResponseWriter, r *http.Request) {
	const perm = "account:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Unauthorized)
		return
	}

	result, err := controller.svc.FetchAccounts()
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *accountController) FetchAccountBySlug(w http.ResponseWriter, r *http.Request) {
	const perm = "account:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Unauthorized)
		return
	}

	slug := mux.Vars(r)["slug"]

	result, err := controller.svc.FetchAccountBySlug(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *accountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	result, err := controller.svc.CreateAccount(account)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *accountController) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.ServerError)
		return
	}

	slug := mux.Vars(r)["slug"]
	account.Slug = slug

	result, err := controller.svc.UpdateAccount(account)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *accountController) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	const perm = "account:delete"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Unauthorized)
		return
	}
	slug := mux.Vars(r)["slug"]

	err = controller.svc.DeleteAccount(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnEmptyBody(w, http.StatusNoContent)
}

func (controller *accountController) Login(w http.ResponseWriter, r *http.Request) {
	var auth model.Auth
	err := json.NewDecoder(r.Body).Decode(&auth)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.ServerError)
		return
	}

	token, err := controller.svc.SignIn(auth)

	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    token,
		MaxAge:   57600,
		HttpOnly: true,
	})

	helper.ReturnSuccess(w, struct{ Token string }{token})
}
