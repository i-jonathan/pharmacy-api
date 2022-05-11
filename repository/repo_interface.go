package repository

import "github.com/i-jonathan/pharmacy-api/model"

// account management repositories

type PermissionRepository interface {
	FetchPermissions() ([]model.Permission, error)
	FetchPermissionByID(int) (model.Permission, error)
	CreatePermission(model.Permission) (int, error)
	UpdatePermission(model.Permission) error
	DeletePermission(int) error
}

type RoleRepository interface {
	FetchRoles() ([]model.Role, error)
	FetchRoleByID(int) (model.Role, error)
	CreateRole(model.Role) (int, error)
	UpdateRole(model.Role) error
	DeleteRole(int) error
}

type AccountRepository interface {
	FetchAccounts() ([]model.Account, error)
	FetchAccountByID(int) (model.Account, error)
	CreateAccount(model.Account) (int, error)
	UpdateAccount(model.Account) error
	DeleteAccount(int) error
}

// inventory management repositories

type CategoryRepository interface {
	FetchCategories() ([]model.Category, error)
	FetchCategoryByID(int) (model.Category, error)
	CreateCategory(model.Category) (int, error)
	UpdateCategory(model.Category) error
	DeleteCategory(int) error
}

type SupplierRepository interface {
	FetchSuppliers() ([]model.Supplier, error)
	FetchSupplierByID(int) (model.Supplier, error)
	CreateSupplier(model.Supplier) (int, error)
	UpdateSupplier(model.Supplier) error
	DeleteSupplier(int) error
}
type ProductRepository interface {
	FetchProducts() ([]model.Product, error)
	FetchProductByID(int) (model.Product, error)
	CreateProduct(model.Product) (int, error)
	UpdateProduct(model.Product) error
	DeleteProduct(int) error
}

// sales management repositories

type PaymentMethodRepository interface {
	FetchPaymentMethods() ([]model.PaymentMethod, error)
	FetchPaymentMethodByID(int) (model.PaymentMethod, error)
	CreatePaymentMethod(model.PaymentMethod) (int, error)
	UpdatePaymentMethod(model.PaymentMethod) error
	DeletePaymentMethod(int) error
}
type OrderItemRepository interface {
	FetchOrderItems() ([]model.OrderItem, error)
	FetchOrderItemByID(int) (model.OrderItem, error)
	CreateOrderItem(model.OrderItem) (int, error)
	UpdateOrderItem(model.OrderItem) error
	DeleteOrderItem(int) error
}
type OrderRepository interface {
	FetchOrders() ([]model.Order, error)
	FetchOrderByID(int) (model.Order, error)
	CreateOrder(model.Order) (int, error)
	UpdateOrder(model.Order) error
	DeleteOrder(int) error
}

type ReturnRepository interface {
	FetchReturns() ([]model.Return, error)
	FetchReturnByID(int) (model.Return, error)
	CreateReturn(model.Return) (int, error)
	UpdateReturn(model.Return) error
	DeleteReturn(int) error
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
	ReturnRepository
}
