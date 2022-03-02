package service

import (
	"wblvl0/internal/cache"
	"wblvl0/internal/model"
	"wblvl0/internal/repository"
)

type OrderService struct {
	repository repository.Order
	cache      *cache.Cache
}

func NewOrderService(repository repository.Order, cache *cache.Cache) *OrderService {
	return &OrderService{repository: repository, cache: cache}
}

func (r *OrderService) CreateOrder(order model.Order) (string, error) {
	return "", nil
}

func (r *OrderService) GetOrder(uid string) (model.Order, error) {
	var order model.Order
	order, found := r.cache.Get(uid)
	if !found {
		pgOrder, err := r.repository.GetOrder(uid)
		if err != nil {
			return model.Order{}, err
		}
		r.cache.Set(order.UID, order)
		return pgOrder, nil
	}
	return order, nil
}
