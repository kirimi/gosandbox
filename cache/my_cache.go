package cache

import "sync"

type MyCache struct {
	storage map[string]string
	mu      *sync.Mutex
}

func NewMyCache() Cache {
	return &MyCache{
		storage: make(map[string]string),
		mu:      &sync.Mutex{},
	}
}

func (c *MyCache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value

	return nil
}

func (c *MyCache) Get(key string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.storage[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func (c *MyCache) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.storage, key)

	return nil
}
