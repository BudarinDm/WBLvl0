package cache

import (
	"sync"
	"time"
	"wblvl0/internal/model"
)

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
}

type Item struct {
	Value      model.Order
	Created    time.Time
	Expiration int64
}

func NewCache(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)

	cache := Cache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	if cleanupInterval > 0 {
		cache.StartGC()
	}

	return &cache
}

func (c *Cache) Set(key string, value model.Order) {

	var expiration int64
	expiration = time.Now().Add(c.defaultExpiration).UnixNano()

	c.Lock()

	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}

}

func (c *Cache) Get(key string) (model.Order, bool) {

	c.RLock()

	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return model.Order{}, false
	}

	if time.Now().UnixNano() > item.Expiration {
		return model.Order{}, false
	}

	return item.Value, true
}

func (c *Cache) StartGC() {
	go c.GC()
}

func (c *Cache) GC() {

	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		// Ищем элементы с истекшим временем жизни и удаляем из хранилища
		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}

	}

}

func (c *Cache) expiredKeys() (keys []string) {

	c.RLock()

	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration {
			keys = append(keys, k)
		}
	}

	return
}

func (c *Cache) clearItems(keys []string) {

	c.Lock()

	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
