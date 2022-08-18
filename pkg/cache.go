package pkg

import (
	"sync"
	"time"

	"github.com/energywork/simple-cache/internal"
)

type Cache struct {
	locker  *sync.RWMutex
	storage map[string]internal.Item
}

func New() *Cache {
	return &Cache{
		locker:  new(sync.RWMutex),
		storage: make(map[string]internal.Item),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	if c.storage == nil {
		return //nothing will happen
	}

	c.locker.Lock()
	c.storage[key] = internal.Item{
		ExpirationTime: internal.GetExpirationTime(ttl),
		Value:          value,
	}
	c.locker.Unlock()
}

func (c *Cache) Get(key string) (interface{}, error) {
	if c.storage == nil {
		return nil, internal.ErrNotInitialized
	}

	c.locker.Lock()
	val, ok := c.storage[key]
	if !ok {
		return nil, internal.ErrNotFound
	}
	c.locker.Unlock()

	if val.Expired() {
		return nil, internal.ErrExpired
	}

	return val, nil
}

func (c *Cache) Delete(key string) {
	if c.storage == nil {
		return
	}

	c.locker.Lock()
	delete(c.storage, key)
	c.locker.Unlock()
}

func (c *Cache) Clear() {
	// todo: ...
}
