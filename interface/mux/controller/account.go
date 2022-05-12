package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/interface/mux/helper"
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/service"
	"log"
	"net/http"
)

type accountController struct {
	svc service.AccountUseCase
}

func NewAccountController(s service.AccountUseCase) *accountController {
	return &accountController{s}
}

func (controller *accountController) FetchAccounts(w http.ResponseWriter, r *http.Request) {
	result, err := controller.svc.FetchAccounts()
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *accountController) FetchAccountBySlug(w http.ResponseWriter, r *http.Request) {
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
		helper.ReturnFailure(w, appError.ServerError)
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
	slug := mux.Vars(r)["slug"]

	err := controller.svc.DeleteAccount(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
	}

	helper.ReturnDelete(w)
}
