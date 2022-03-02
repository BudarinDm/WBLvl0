package cache

import (
	"reflect"
	"wblvl0/internal/model"
	"wblvl0/internal/service"
)

type Cache struct {
	Orders map[string]model.Order
}

func NewCache(orders map[string]model.Order) *Cache {
	return &Cache{Orders: orders}
}

func (c *Cache) Add(order model.Order) {
	c.Orders[order.UID] = order
}

func (c *Cache) GetInHandler(uid string, service *service.Service) (model.Order, error) {
	order := c.Orders[uid]
	if reflect.ValueOf(order).IsZero() {
		pqOrder, err := service.Order.GetOrder(uid)
		c.Add(pqOrder)
		if err != nil {
			return model.Order{}, err
		}
		return pqOrder, nil
	}
	return order, nil
}
