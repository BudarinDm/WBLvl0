package service

import (
	"wblvl0/internal/model"
	"wblvl0/internal/repository"
)

type Order interface {
	CreateOrder(order model.Order) (string, error)
	GetOrder(uid string) (model.Order, error)
}

type Service struct {
	Order
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repository),
	}
}
