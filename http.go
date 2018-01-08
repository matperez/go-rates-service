package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type rateRequest struct {
	Currency string `json:"currency"`
}

type rateResponse struct {
	Value float64 `json:"value"`
	Error string  `json:"error,omitempty"`
}

func endcodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeRateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request rateRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func makeGetRateEndpoint(svc RatesService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(rateRequest)
		v, err := svc.GetRate(ctx, req.Currency)
		if err != nil {
			return rateResponse{v, err.Error()}, nil
		}
		return rateResponse{v, ""}, nil
	}
}
