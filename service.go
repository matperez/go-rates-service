package main

import (
	"context"
	"time"

	"github.com/matperez/cbr-http-service/cache"
	"github.com/matperez/go-cbr-client"
)

// RatesService is the service itself
type RatesService interface {
	GetRate(context.Context, string) (float64, error)
}

type ratesService struct {
	client cbr.Client
	cache  cache.Cache
}

func (s ratesService) GetRate(_ context.Context, currency string) (float64, error) {
	return s.cache.GetOrSet(currency, func(currency string) (float64, error) {
		rate, err := s.client.GetRate(currency, time.Now())
		if err != nil {
			return 0, err
		}
		return rate, nil
	})
}

// NewService is a fabric function
func NewService(client cbr.Client, cache cache.Cache) RatesService {
	return ratesService{
		client,
		cache,
	}
}
