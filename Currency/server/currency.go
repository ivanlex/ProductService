package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/kevin/currency/protos/currency"
)

type Currency struct {
	log hclog.Logger
}

func NewCurrency(l hclog.Logger) *Currency {
	return &Currency{l}
}

func (c *Currency) GetRate(ctx context.Context, rr *currency.RateRequest) (*currency.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())

	return &currency.RateResponse{Rate: 0.5}, nil
}
