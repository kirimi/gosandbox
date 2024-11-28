package cache

import "errors"

var ErrNotFound = errors.New("value  not found")

type Cache interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(ket string) error
}

type CacheWithMetrics interface {
	Cache
	Metrics
}
