package service

import (
	"log"

	appError "github.com/i-jonathan/pharmacy-api/error"
	"github.com/i-jonathan/pharmacy-api/model"
)

func (service *inventoryService) FetchSuppliers() ([]model.Supplier, error) {
	result, err := service.repo.FetchSuppliers()
	if err != nil {
		log.Println(err)
		return nil, appError.ServerError
	}
	return result, nil
}

func (service *inventoryService) FetchSupplierBySlug(slug string) (model.Supplier, error) {
	id, err := model.DecodeSlugToID(slug)
	if err != nil {
		log.Println(err)
		return model.Supplier{}, appError.BadRequest
	}

	result, err := service.repo.FetchSupplierByID(id)
	if err != nil {
		if err == appError.NotFound {
			return model.Supplier{}, err
		}
		return model.Supplier{}, appError.ServerError
	}

	return result, nil
}

func (service *inventoryService) CreateSupplier(supplier model.Supplier) (model.Supplier, error) {
	if !supplier.Valid() {
		return model.Supplier{}, appError.BadRequest
	}

	result, err := service.repo.CreateSupplier(supplier)
	if err != nil {
		log.Println(err)
		return model.Supplier{}, appError.ServerError
	}

	supplier.ID = result
	supplier.Slug, err = model.EncodeIDToSlug(supplier.ID)
	if err != nil {
		log.Println(err)
	}
	return supplier, nil
}

func (service *inventoryService) UpdateSupplier(supplier model.Supplier) (model.Supplier, error) {
	id, err := model.DecodeSlugToID(supplier.Slug)
	if err != nil {
		log.Println(err)
		return model.Supplier{}, appError.BadRequest
	}

	if !supplier.Valid() {
		return model.Supplier{}, appError.BadRequest
	}

	_, err = service.repo.FetchSupplierByID(id)
	if err != nil {
		if err == appError.NotFound {
			return model.Supplier{}, err
		}
		log.Println(err)
		return model.Supplier{}, appError.ServerError
	}

	err = service.repo.UpdateSupplier(supplier)
	if err != nil {
		log.Println(err)
		return model.Supplier{}, appError.ServerError
	}

	return supplier, nil

}

func (service *inventoryService) DeleteSupplier(slug string) error {
	id, err := model.DecodeSlugToID(slug)
	if err != nil {
		log.Println(err)
		return appError.BadRequest
	}

	_, err = service.repo.FetchSupplierByID(id)
	if err != nil {
		if err == appError.NotFound {
			return err
		}
		log.Println(err)
		return appError.ServerError
	}

	err = service.repo.DeleteSupplier(id)
	if err != nil {
		log.Println(err)
		return appError.ServerError
	}

	return nil
}
