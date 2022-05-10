package service

import (
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/repository"
)

type inventoryService struct {
	repo repository.PharmacyRepository
}

func NewInventoryService(r repository.PharmacyRepository) *inventoryService {
	return &inventoryService{r}
}

// Categories

func (i *inventoryService) FetchCategories() ([]model.Category, error) {
	panic("implement me")
}

func (i *inventoryService) FetchCategoryBySlug(slug string) (model.Category, error) {
	panic("implement me")
}

func (i *inventoryService) CreateCategory(category model.Category) (model.Category, error) {
	panic("implement me")
}

func (i *inventoryService) UpdateCategory(category model.Category) (model.Category, error) {
	panic("implement me")
}

func (i *inventoryService) DeleteCategory(slug string) error {
	panic("implement me")
}

// Supplier

func (i *inventoryService) FetchSuppliers() ([]model.Supplier, error) {
	panic("implement me")
}

func (i *inventoryService) FetchSupplierBySlug(slug string) (model.Supplier, error) {
	panic("implement me")
}

func (i *inventoryService) CreateSupplier(supplier model.Supplier) (model.Supplier, error) {
	panic("implement me")
}

func (i *inventoryService) UpdateSupplier(supplier model.Supplier) (model.Supplier, error) {
	panic("implement me")
}

func (i *inventoryService) DeleteSupplier(slug string) error {
	panic("implement me")
}

// Products

func (i *inventoryService) FetchProducts() ([]model.Product, error) {
	panic("implement me")
}

func (i *inventoryService) FetchProductBySlug(slug string) (model.Product, error) {
	panic("implement me")
}

func (i *inventoryService) CreateProduct(product model.Product) (model.Product, error) {
	panic("implement me")
}

func (i *inventoryService) UpdateProduct(product model.Product) (model.Product, error) {
	panic("implement me")
}

func (i *inventoryService) DeleteProduct(slug string) error {
	panic("implement me")
}
