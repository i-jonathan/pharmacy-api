package db

import "github.com/i-jonathan/pharmacy-api/model"

func (r *repo) FetchCategories() ([]model.Category, error) {
	panic("implement me")
}

func (r *repo) FetchCategoryByID(i int) (model.Category, error) {
	panic("implement me")
}

func (r *repo) CreateCategory(category model.Category) (int, error) {
	panic("implement me")
}

func (r *repo) UpdateCategory(category model.Category) error {
	panic("implement me")
}

func (r *repo) DeleteCategory(i int) error {
	panic("implement me")
}

func (r *repo) FetchSuppliers() ([]model.Supplier, error) {
	panic("implement me")
}

func (r *repo) FetchSupplierByID(i int) (model.Supplier, error) {
	panic("implement me")
}

func (r *repo) CreateSupplier(supplier model.Supplier) (int, error) {
	panic("implement me")
}

func (r *repo) UpdateSupplier(supplier model.Supplier) error {
	panic("implement me")
}

func (r *repo) DeleteSupplier(i int) error {
	panic("implement me")
}

func (r *repo) FetchProducts() ([]model.Product, error) {
	panic("implement me")
}

func (r *repo) FetchProductByID(i int) (model.Product, error) {
	panic("implement me")
}

func (r *repo) CreateProduct(product model.Product) (int, error) {
	panic("implement me")
}

func (r *repo) UpdateProduct(product model.Product) error {
	panic("implement me")
}

func (r *repo) DeleteProduct(i int) error {
	panic("implement me")
}
