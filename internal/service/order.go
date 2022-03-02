package service

import (
	"wblvl0/internal/model"
	"wblvl0/internal/repository"
)

type OrderService struct {
	repository repository.Order
}

func NewOrderService(repository repository.Order) *OrderService {
	return &OrderService{repository: repository}
}

func (r *OrderService) CreateOrder(order model.Order) (string, error) {
	return "", nil
}

func (r *OrderService) GetOrder(uid string) (model.Order, error) {
	return r.repository.GetOrder(uid)
}
