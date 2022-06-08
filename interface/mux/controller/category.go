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

type categoryController struct {
	svc service.CategoryUseCase
}

func NewCategoryController(s service.CategoryUseCase) *categoryController {
	return &categoryController{s}
}

func (controller *categoryController) FetchCategories(w http.ResponseWriter, r *http.Request) {
	const perm = "category:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	result, err := controller.svc.FetchCategories()
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *categoryController) FetchCategoryBySlug(w http.ResponseWriter, r *http.Request) {
	const perm = "category:read"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	slug := mux.Vars(r)["slug"]

	result, err := controller.svc.FetchCategoryBySlug(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *categoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	const perm = "category:create"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	var category model.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, appError.BadRequest)
		return
	}

	result, err := controller.svc.CreateCategory(category)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *categoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	const perm = "category:update"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	var category model.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		log.Println(err)
		helper.ReturnFailure(w, err)
		return
	}

	category.Slug = mux.Vars(r)["slug"]

	result, err := controller.svc.UpdateCategory(category)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnSuccess(w, result)
}

func (controller *categoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	const perm = "category:delete"
	allowed, err := model.CheckPermission(perm, r)
	if !allowed {
		log.Println(err)
		helper.ReturnFailure(w, appError.Forbidden)
		return
	}

	slug := mux.Vars(r)["slug"]
	err = controller.svc.DeleteCategory(slug)
	if err != nil {
		helper.ReturnFailure(w, err)
		return
	}

	helper.ReturnEmptyBody(w, http.StatusNoContent)
}
