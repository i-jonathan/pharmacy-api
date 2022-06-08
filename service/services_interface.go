package service

import "github.com/i-jonathan/pharmacy-api/model"

// account management repositories

type AuthUseCase interface {
	Logout(hash, token string) error
}

type PermissionUseCase interface {
	FetchPermissions() ([]model.Permission, error)
	FetchPermissionBySlug(string) (model.Permission, error)
	CreatePermission(model.Permission) (model.Permission, error)
	UpdatePermission(model.Permission) (model.Permission, error)
	DeletePermission(string) error
}

type RoleUseCase interface {
	FetchRoles() ([]model.Role, error)
	FetchRoleBySlug(string) (model.Role, error)
	CreateRole(model.Role) (model.Role, error)
	UpdateRole(model.Role) error
	DeleteRole(string) error
}

type AccountUseCase interface {
	SignIn(auth model.Auth) (string, error)
	FetchAccounts() ([]model.Account, error)
	FetchAccountBySlug(string) (model.Account, error)
	CreateAccount(model.Account) (model.Account, error)
	UpdateAccount(model.Account) (model.Account, error)
	DeleteAccount(string) error
}

// inventory management repositories

type CategoryUseCase interface {
	FetchCategories() ([]model.Category, error)
	FetchCategoryBySlug(string) (model.Category, error)
	CreateCategory(model.Category) (model.Category, error)
	UpdateCategory(model.Category) (model.Category, error)
	DeleteCategory(string) error
}

type SupplierUseCase interface {
	FetchSuppliers() ([]model.Supplier, error)
	FetchSupplierBySlug(string) (model.Supplier, error)
	CreateSupplier(model.Supplier) (model.Supplier, error)
	UpdateSupplier(model.Supplier) (model.Supplier, error)
	DeleteSupplier(string) error
}
type ProductUseCase interface {
	FetchProducts() ([]model.Product, error)
	FetchProductBySlug(string) (model.Product, error)
	CreateProduct(model.Product) (model.Product, error)
	UpdateProduct(model.Product) (model.Product, error)
	DeleteProduct(string) error
}

type InventoryUseCase interface {
	CategoryUseCase
	ProductUseCase
	SupplierUseCase
}

// sales management repositories

type PaymentMethodUseCase interface {
	FetchPaymentMethods() ([]model.PaymentMethod, error)
	FetchPaymentMethodBySlug(string) (model.PaymentMethod, error)
	CreatePaymentMethod(model.PaymentMethod) (model.PaymentMethod, error)
	UpdatePaymentMethod(model.PaymentMethod) (model.PaymentMethod, error)
	DeletePaymentMethod(string) error
}
type OrderItemUseCase interface {
	FetchOrderItems() ([]model.OrderItem, error)
	FetchOrderItemBySlug(string) (model.OrderItem, error)
	CreateOrderItem(model.OrderItem) (model.OrderItem, error)
	UpdateOrderItem(model.OrderItem) (model.OrderItem, error)
	DeleteOrderItem(string) error
}
type OrderUseCase interface {
	FetchOrders() ([]model.Order, error)
	FetchOrderBySlug(string) (model.Order, error)
	CreateOrder(model.Order) (model.Order, error)
	UpdateOrder(model.Order) (model.Order, error)
	DeleteOrder(string) error
}

type ReturnUseCase interface {
	FetchReturns() ([]model.Return, error)
	FetchReturnBySlug(string) (model.Return, error)
	CreateReturn(model.Return) (model.Return, error)
	UpdateReturn(model.Return) (model.Return, error)
	DeleteReturn(string) error
}

type PharmacyUseCase interface {
	PermissionUseCase
	RoleUseCase
	AccountUseCase
	CategoryUseCase
	SupplierUseCase
	ProductUseCase
	PaymentMethodUseCase
	OrderItemUseCase
	OrderUseCase
	ReturnUseCase
}
