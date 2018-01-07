package main

import (
	"context"
	"time"

	"github.com/matperez/go-cbr-client"
)

// RatesService is the service itself
type RatesService interface {
	GetRate(context.Context, string) (float64, error)
}

type ratesService struct {
	client cbr.Client
}

func (s ratesService) GetRate(_ context.Context, currecy string) (float64, error) {
	rate, err := s.client.GetRate(currecy, time.Now())
	if err != nil {
		return 0, err
	}
	return rate, nil
}

// NewService is a fabric function
func NewService(client cbr.Client) RatesService {
	return ratesService{client}
}
