package router

import (
	"github.com/i-jonathan/pharmacy-api/interface/mux/controller"
	"github.com/i-jonathan/pharmacy-api/service"
	"net/http"
)

func InitInventoryRouter(svc service.InventoryUseCase) {
	productRoutes(svc)
	categoryRoutes(svc)
	supplierRoutes(svc)
}

func productRoutes(svc service.ProductUseCase) {
	productController := controller.NewProductController(svc)
	productRouter := router.PathPrefix("/product").Subrouter()

	productRouter.HandleFunc("", productController.CreateProduct).Methods(http.MethodPost)
	productRouter.HandleFunc("", productController.FetchProducts).Methods(http.MethodGet)
	productRouter.HandleFunc("/{slug}", productController.FetchProductBySlug).Methods(http.MethodGet)
	productRouter.HandleFunc("/{slug}", productController.UpdateProduct).Methods(http.MethodPut)
	productRouter.HandleFunc("/{slug}", productController.DeleteProduct).Methods(http.MethodDelete)
}

func categoryRoutes(svc service.CategoryUseCase) {
	categoryController := controller.NewCategoryController(svc)
	categoryRouter := router.PathPrefix("/category").Subrouter()

	categoryRouter.HandleFunc("", categoryController.CreateCategory).Methods(http.MethodPost)
	categoryRouter.HandleFunc("", categoryController.FetchCategories).Methods(http.MethodGet)
	categoryRouter.HandleFunc("/{slug}", categoryController.FetchCategoryBySlug).Methods(http.MethodGet)
	categoryRouter.HandleFunc("/{slug}", categoryController.UpdateCategory).Methods(http.MethodPut)
	categoryRouter.HandleFunc("/{slug}", categoryController.DeleteCategory).Methods(http.MethodDelete)
}

func supplierRoutes(svc service.SupplierUseCase) {
	supplierController := controller.NewSupplierController(svc)
	supplierRouter := router.PathPrefix("/supplier").Subrouter()

	supplierRouter.HandleFunc("", supplierController.CreateSupplier).Methods(http.MethodPost)
	supplierRouter.HandleFunc("", supplierController.FetchSuppliers).Methods(http.MethodGet)
	supplierRouter.HandleFunc("/{slug}", supplierController.FetchSupplierBySlug).Methods(http.MethodGet)
	supplierRouter.HandleFunc("/{slug}", supplierController.UpdateSupplier).Methods(http.MethodPut)
	supplierRouter.HandleFunc("/{slug}", supplierController.DeleteSupplier).Methods(http.MethodDelete)
}
