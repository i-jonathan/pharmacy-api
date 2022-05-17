package db

import "github.com/i-jonathan/pharmacy-api/model"

func (r *repo) FetchPaymentMethods() ([]model.PaymentMethod, error) {
	panic("implement me")
}

func (r *repo) FetchPaymentMethodByID(i int) (model.PaymentMethod, error) {
	panic("implement me")
}

func (r *repo) CreatePaymentMethod(method model.PaymentMethod) (int, error) {
	panic("implement me")
}

func (r *repo) UpdatePaymentMethod(method model.PaymentMethod) error {
	panic("implement me")
}

func (r *repo) DeletePaymentMethod(i int) error {
	panic("implement me")
}

func (r *repo) FetchOrderItems() ([]model.OrderItem, error) {
	panic("implement me")
}

func (r *repo) FetchOrderItemByID(i int) (model.OrderItem, error) {
	panic("implement me")
}

func (r *repo) CreateOrderItem(item model.OrderItem) (int, error) {
	panic("implement me")
}

func (r *repo) UpdateOrderItem(item model.OrderItem) error {
	panic("implement me")
}

func (r *repo) DeleteOrderItem(i int) error {
	panic("implement me")
}

func (r *repo) FetchOrders() ([]model.Order, error) {
	panic("implement me")
}

func (r *repo) FetchOrderByID(i int) (model.Order, error) {
	panic("implement me")
}

func (r *repo) CreateOrder(order model.Order) (int, error) {
	panic("implement me")
}

func (r *repo) UpdateOrder(order model.Order) error {
	panic("implement me")
}

func (r *repo) DeleteOrder(i int) error {
	panic("implement me")
}

func (r *repo) FetchReturns() ([]model.Return, error) {
	panic("implement me")
}

func (r *repo) FetchReturnByID(i int) (model.Return, error) {
	panic("implement me")
}

func (r *repo) CreateReturn(m model.Return) (int, error) {
	panic("implement me")
}

func (r *repo) UpdateReturn(m model.Return) error {
	panic("implement me")
}

func (r *repo) DeleteReturn(i int) error {
	panic("implement me")
}
