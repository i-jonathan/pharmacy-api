package service

import (
	"github.com/i-jonathan/pharmacy-api/model"
	"github.com/i-jonathan/pharmacy-api/repository"
)

type salesService struct {
	repo repository.PharmacyRepository
}

// Payment Method

func (s *salesService) FetchPaymentMethods() ([]model.PaymentMethod, error) {
	panic("implement me")
}

func (s *salesService) FetchPaymentMethodBySlug(slug string) (model.PaymentMethod, error) {
	panic("implement me")
}

func (s *salesService) CreatePaymentMethod(method model.PaymentMethod) (model.PaymentMethod, error) {
	panic("implement me")
}

func (s *salesService) UpdatePaymentMethod(method model.PaymentMethod) (model.PaymentMethod, error) {
	panic("implement me")
}

func (s *salesService) DeletePaymentMethod(slug string) error {
	panic("implement me")
}

// Order Item

func (s *salesService) FetchOrderItems() ([]model.OrderItem, error) {
	panic("implement me")
}

func (s *salesService) FetchOrderItemBySlug(slug string) (model.OrderItem, error) {
	panic("implement me")
}

func (s *salesService) CreateOrderItem(item model.OrderItem) (model.OrderItem, error) {
	panic("implement me")
}

func (s *salesService) UpdateOrderItem(item model.OrderItem) (model.OrderItem, error) {
	panic("implement me")
}

func (s *salesService) DeleteOrderItem(slug string) error {
	panic("implement me")
}

// Order

func (s *salesService) FetchOrders() ([]model.Order, error) {
	panic("implement me")
}

func (s *salesService) FetchOrderBySlug(slug string) (model.Order, error) {
	panic("implement me")
}

func (s *salesService) CreateOrder(order model.Order) (model.Order, error) {
	panic("implement me")
}

func (s *salesService) UpdateOrder(order model.Order) (model.Order, error) {
	panic("implement me")
}

func (s *salesService) DeleteOrder(slug string) error {
	panic("implement me")
}

// Return

func (s *salesService) FetchReturns() ([]model.Return, error) {
	panic("implement me")
}

func (s *salesService) FetchReturnBySlug(slug string) (model.Return, error) {
	panic("implement me")
}

func (s *salesService) CreateReturn(r model.Return) (model.Return, error) {
	panic("implement me")
}

func (s *salesService) UpdateReturn(r model.Return) (model.Return, error) {
	panic("implement me")
}

func (s *salesService) DeleteReturn(slug string) error {
	panic("implement me")
}
