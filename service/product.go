package service

import (
	"log"

	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
)

func (service *inventoryService) FetchProducts() ([]model.Product, error) {
	result, err := service.repo.FetchProducts()

	if err != nil {
		log.Println(err)
		return nil, appError.ServerError
	}

	return result, nil
}

func (service *inventoryService) FetchProductBySlug(slug string) (model.Product, error) {
	id, err := model.DecodeSlugToID(slug)
	if err != nil {
		log.Println(err)
		return model.Product{}, appError.BadRequest
	}

	result, err := service.repo.FetchProductByID(id)
	if err != nil {
		if err == appError.NotFound {
			return model.Product{}, err
		}
		log.Println(err)
		return model.Product{}, appError.ServerError
	}

	return result, nil
}

func (service *inventoryService) CreateProduct(product model.Product) (model.Product, error) {
	if !product.Valid() {
		return model.Product{}, appError.BadRequest
	}

	result, err := service.repo.CreateProduct(product)

	if err != nil {
		log.Println(err)
		return model.Product{}, appError.ServerError
	}

	product.ID = result
	product.Slug, err = model.EncodeIDToSlug(product.ID)
	if err != nil {
		log.Println(err)
	}
	return product, nil
}

func (service *inventoryService) UpdateProduct(product model.Product) (model.Product, error) {
	id, err := model.DecodeSlugToID(product.Slug)
	if err != nil {
		log.Println(err)
		return model.Product{}, appError.BadRequest
	}

	product.ID = id
	_, err = service.repo.FetchProductByID(product.ID)
	if err != nil {
		if err == appError.NotFound {
			return model.Product{}, err
		}
		log.Println(err)
		return model.Product{}, appError.ServerError
	}

	if !product.Valid() {
		return model.Product{}, appError.BadRequest
	}

	err = service.repo.UpdateProduct(product)
	if err != nil {
		log.Println(err)
		return model.Product{}, appError.ServerError
	}

	return product, nil
}

func (service *inventoryService) DeleteProduct(slug string) error {
	id, err := model.DecodeSlugToID(slug)
	if err != nil {
		log.Println(err)
		return appError.BadRequest
	}

	_, err = service.repo.FetchProductByID(id)
	if err != nil {
		if err == appError.NotFound {
			return err
		}
		log.Println(err)
		return appError.ServerError
	}

	err = service.repo.DeleteProduct(id)
	if err != nil {
		log.Println(err)
		return appError.ServerError
	}

	return nil
}
