package repository

import "github.com/i-jonathan/pharmacy-api/model"

// account management repositories

type PermissionRepository interface {
	FetchPermissions() ([]model.Permission, error)
	FetchPermissionBySlug(string) (model.Permission, error)
	CreatePermission(model.Permission) (int, error)
	UpdatePermission(model.Permission) error
	DeletePermission(string) error
}

type RoleRepository interface {
	FetchRoles() ([]model.Role, error)
	FetchRoleBySlug(string) (model.Role, error)
	CreateRole(model.Role) (int, error)
	UpdateRole(model.Role) error
	DeleteRole(string) error
}

type AccountRepository interface {
	FetchAccounts() ([]model.Account, error)
	FetchAccountBySlug(string) (model.Account, error)
	CreateAccount(model.Account) (int, error)
	UpdateAccount(model.Account) error
	DeleteAccount(string) error
}

// inventory management repositories

type CategoryRepository interface {
	FetchCategories() ([]model.Category, error)
	FetchCategoryBySlug(string) (model.Category, error)
	CreateCategory(model.Category) (int, error)
	UpdateCategory(model.Category) error
	DeleteCategory(string) error
}

type SupplierRepository interface {
	FetchSuppliers() ([]model.Supplier, error)
	FetchSupplierBySlug(string) (model.Supplier, error)
	CreateSupplier(model.Supplier) (int, error)
	UpdateSupplier(model.Supplier) error
	DeleteSupplier(string) error
}
type ProductRepository interface {
	FetchProducts() ([]model.Product, error)
	FetchProductBySlug(string) (model.Product, error)
	CreateProduct(model.Product) (int, error)
	UpdateProduct(model.Product) error
	DeleteProduct(string) error
}

// sales management repositories

type PaymentMethodRepository interface {
	FetchPaymentMethods() ([]model.PaymentMethod, error)
	FetchPaymentMethodBySlug(string) (model.PaymentMethod, error)
	CreatePaymentMethod(model.PaymentMethod) (int, error)
	UpdatePaymentMethod(model.PaymentMethod) error
	DeletePaymentMethod(string) error
}
type OrderItemRepository interface {
	FetchOrderItems() ([]model.OrderItem, error)
	FetchOrderItemBySlug(string) (model.OrderItem, error)
	CreateOrderItem(model.OrderItem) (int, error)
	UpdateOrderItem(model.OrderItem) error
	DeleteOrderItem(string) error
}
type OrderRepository interface {
	FetchOrders() ([]model.Order, error)
	FetchOrderBySlug(string) (model.Order, error)
	CreateOrder(model.Order) (int, error)
	UpdateOrder(model.Order) error
	DeleteOrder(string) error
}

type PharmacyRepository interface {
	PermissionRepository
	RoleRepository
	AccountRepository
	CategoryRepository
	SupplierRepository
	ProductRepository
	PaymentMethodRepository
	OrderItemRepository
	OrderRepository
}
