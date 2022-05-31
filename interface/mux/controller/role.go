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

type roleController struct {
	svc service.RoleUseCase
}

func NewRoleController(s service.RoleUseCase) *roleController {
	return &roleController{s}
}

func (controller *roleController) FetchRoles(w http.ResponseWriter, r *http.Request) {
	result, err := controller.svc.FetchRoles()
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *roleController) FetchRoleBySlug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]
	result, err := controller.svc.FetchRoleBySlug(slug)

	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *roleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role model.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	result, err := controller.svc.CreateRole(role)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *roleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	var role model.Role
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, err)
		return
	}

	role.Slug = mux.Vars(r)["slug"]

	err = controller.svc.UpdateRole(role)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, role)
}

func (controller *roleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	err := controller.svc.DeleteRole(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnEmptyBody(w, http.StatusNoContent)
}