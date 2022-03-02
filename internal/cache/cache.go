package cache

import "wblvl0/internal/model"

type Cache struct {
	Orders map[string]model.Order
}

func (c *Cache) Add(order model.Order) {
	c.Orders[order.UID] = order
}
