package service

import (
	"log"

	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
)

func (service *inventoryService) FetchCategories() ([]model.Category, error) {
	result, err := service.repo.FetchCategories()
	if err != nil {
		log.Println(err)
		return nil, appError.ServerError
	}

	return result, nil
}

func (service *inventoryService) FetchCategoryBySlug(slug string) (model.Category, error) {
	id, err := model.DecodeSlugToID(slug)
	if err != nil {
		log.Println(err)
		return model.Category{}, appError.BadRequest
	}

	result, err := service.repo.FetchCategoryByID(id)
	if err != nil {
		if err == appError.NotFound {
			return model.Category{}, err
		}
		return model.Category{}, appError.ServerError
	}

	return result, nil
}

func (service *inventoryService) CreateCategory(category model.Category) (model.Category, error) {
	valid := category.Valid()

	if !valid {
		return model.Category{}, appError.BadRequest
	}

	result, err := service.repo.CreateCategory(category)

	if err != nil {
		log.Println(err)
		return model.Category{}, appError.ServerError
	}

	category.ID = result
	category.Slug, err = model.EncodeIDToSlug(category.ID)
	if err != nil {
		log.Println(err)
	}

	return category, nil
}

func (service *inventoryService) UpdateCategory(category model.Category) (model.Category, error) {
	id, err := model.DecodeSlugToID(category.Slug)
	if err != nil {
		log.Println(err)
		return model.Category{}, appError.BadRequest
	}

	category.ID = id
	_, err = service.repo.FetchCategoryByID(category.ID)
	if err != nil {
		if err == appError.NotFound {
			log.Println(err)
			return model.Category{}, err
		}
		return model.Category{}, appError.ServerError
	}

	if !category.Valid() {
		return model.Category{}, appError.BadRequest
	}

	err = service.repo.UpdateCategory(category)
	if err != nil {
		log.Println(err)
		return model.Category{}, appError.ServerError
	}

	return category, nil
}

func (service *inventoryService) DeleteCategory(slug string) error {
	id, err := model.DecodeSlugToID(slug)
	if err != nil {
		log.Println(err)
		return appError.BadRequest
	}

	// Check if category exists
	_, err = service.repo.FetchCategoryByID(id)
	if err != nil {
		log.Println(err)
		if err == appError.NotFound {
			return err
		}
		return appError.ServerError
	}

	err = service.repo.DeleteCategory(id)
	if err != nil {
		log.Println(err)
		return appError.ServerError
	}

	return nil
}
