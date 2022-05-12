package controller

import (
	"github.com/gorilla/mux"
	"github.com/i-jonathan/pharmacy-api/interface/mux/helper"
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
