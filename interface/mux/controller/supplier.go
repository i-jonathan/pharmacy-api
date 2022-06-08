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

type supplierController struct {
	svc service.SupplierUseCase
}

func NewSupplierController(s service.SupplierUseCase) *supplierController {
	return &supplierController{s}
}

func (controller *supplierController) FetchSuppliers(w http.ResponseWriter, r *http.Request) {
	const perm = "supplier:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	result, err := controller.svc.FetchSuppliers()
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *supplierController) FetchSupplierBySlug(w http.ResponseWriter, r *http.Request) {
	const perm = "supplier:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	slug := mux.Vars(r)["slug"]

	result, err := controller.svc.FetchSupplierBySlug(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *supplierController) CreateSupplier(w http.ResponseWriter, r *http.Request) {
	const perm = "supplier:create"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	var supplier model.Supplier
	err = json.NewDecoder(r.Body).Decode(&supplier)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	result, err := controller.svc.CreateSupplier(supplier)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *supplierController) UpdateSupplier(w http.ResponseWriter, r *http.Request) {
	const perm = "supplier:update"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	var supplier model.Supplier
	err = json.NewDecoder(r.Body).Decode(&supplier)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	supplier.Slug = mux.Vars(r)["slug"]

	result, err := controller.svc.UpdateSupplier(supplier)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *supplierController) DeleteSupplier(w http.ResponseWriter, r *http.Request) {
	const perm = "supplier:delete"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	slug := mux.Vars(r)["slug"]
	err = controller.svc.DeleteSupplier(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnEmptyBody(w, http.StatusNoContent)
}
