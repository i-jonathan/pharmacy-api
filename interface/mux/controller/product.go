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

type productController struct {
	svc service.ProductUseCase
}

func NewProductController(s service.ProductUseCase) *productController {
	return &productController{s}
}

func (controller *productController) FetchProducts(w http.ResponseWriter, r *http.Request) {
	const perm = "product:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	result, err := controller.svc.FetchProducts()
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *productController) FetchProductBySlug(w http.ResponseWriter, r *http.Request) {
	const perm = "product:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	slug := mux.Vars(r)["slug"]

	result, err := controller.svc.FetchProductBySlug(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	const perm = "product:create"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	var product model.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	result, err := controller.svc.CreateProduct(product)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *productController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	const perm = "product:update"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	var product model.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	product.Slug = mux.Vars(r)["slug"]

	result, err := controller.svc.UpdateProduct(product)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}
	helper.ReturnSuccess(w, result)
}

func (controller *productController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	const perm = "product:delete"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	slug := mux.Vars(r)["slug"]
	err = controller.svc.DeleteProduct(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnEmptyBody(w, http.StatusNoContent)
}
