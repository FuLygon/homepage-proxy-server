package cache

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, duration time.Duration)
}
type cache struct {
	cache *gocache.Cache
}

func NewCache() Cache {
	return &cache{
		cache: gocache.New(15*time.Minute, 30*time.Minute),
	}
}

// Get retrieves a value from cache
func (c *cache) Get(key string) (interface{}, bool) {
	return c.cache.Get(key)
}

// Set stores a value in cache
func (c *cache) Set(key string, value interface{}, duration time.Duration) {
	c.cache.Set(key, value, duration)
}
