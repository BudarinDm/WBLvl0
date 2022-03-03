package service

import (
	"wblvl0/internal/cache"
	"wblvl0/internal/model"
	"wblvl0/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Order interface {
	CreateOrder(order model.Order) (string, error)
	GetOrder(uid string) (model.Order, error)
}

type Service struct {
	Order
}

func NewService(repository *repository.Repository, cache *cache.Cache) *Service {
	return &Service{
		Order: NewOrderService(repository, cache),
	}
}
