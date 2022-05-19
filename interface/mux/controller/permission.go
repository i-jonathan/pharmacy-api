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

type permissionController struct {
	svc service.PermissionUseCase
}

func NewPermissionController(s service.PermissionUseCase) *permissionController {
	return &permissionController{s}
}

func (controller *permissionController) FetchPermissions(w http.ResponseWriter, _ *http.Request) {
	result, err := controller.svc.FetchPermissions()
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *permissionController) FetchPermissionBySlug(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	result, err := controller.svc.FetchPermissionBySlug(slug)

	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller permissionController) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var permission model.Permission
	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	result, err := controller.svc.CreatePermission(permission)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *permissionController) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	var permission model.Permission
	err := json.NewDecoder(r.Body).Decode(&permission)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	slug := mux.Vars(r)["slug"]
	permission.Slug = slug

	result, err := controller.svc.UpdatePermission(permission)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *permissionController) DeletePermission(w http.ResponseWriter, r *http.Request) {
	slug := mux.Vars(r)["slug"]

	err := controller.svc.DeletePermission(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnEmptyBody(w, http.StatusNoContent)
}
