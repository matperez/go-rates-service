package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// Getter is a getter function
type Getter func(string) (float64, error)

// Cache for currency rates
type Cache interface {
	Get(string) (float64, bool)
	Set(string, float64)
	GetOrSet(string, Getter) (float64, error)
}

type cache struct {
	inner *gocache.Cache
}

// New is a factory method
func New() Cache {
	return cache{gocache.New(5*time.Minute, 10*time.Minute)}
}

func (c cache) GetOrSet(currency string, getter Getter) (float64, error) {
	stored, found := c.Get(currency)
	if found {
		return stored, nil
	}
	fetched, err := getter(currency)
	if err != nil {
		return 0, err
	}
	c.Set(currency, fetched)
	return fetched, nil
}

func (c cache) Get(currency string) (float64, bool) {
	v, found := c.inner.Get(currency)
	if found {
		return v.(float64), true
	}
	return 0, false
}

func (c cache) Set(currency string, value float64) {
	c.inner.Set(currency, value, gocache.DefaultExpiration)
}
