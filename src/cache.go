package src

type Cache struct {
	storage map[string]interface{}
}

func New() *Cache {
	return &Cache{
		storage: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	if c.storage == nil {
		return
	}
	c.storage[key] = value
}

func (c *Cache) Get(key string) interface{} {
	if c.storage == nil {
		return nil
	}
	return c.storage[key] // will be zero value -> nil
}

func (c *Cache) Delete(key string) {
	if c.storage == nil {
		return
	}
	delete(c.storage, key)
}
