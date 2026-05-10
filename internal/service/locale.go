package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type LocaleService struct {
	q *repository.Queries
}

func NewLocaleService(pool *pgxpool.Pool) *LocaleService {
	return &LocaleService{q: repository.New(pool)}
}

type CurrencyResponse struct {
	Code       string  `json:"code"`
	Symbol     string  `json:"symbol"`
	Name       string  `json:"name"`
	RateVsUSD  float64 `json:"rateVsUsd"`
}

type RegionResponse struct {
	Code                string `json:"code"`
	Name                string `json:"name"`
	DefaultCurrencyCode string `json:"defaultCurrencyCode"`
}

func (s *LocaleService) GetCurrencies(ctx context.Context) ([]CurrencyResponse, error) {
	rows, err := s.q.ListCurrencies(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]CurrencyResponse, len(rows))
	for i, r := range rows {
		out[i] = CurrencyResponse{
			Code:      r.Code,
			Symbol:    r.Symbol,
			Name:      r.Name,
			RateVsUSD: numericToFloat(r.RateVsUsd),
		}
	}
	return out, nil
}

func (s *LocaleService) GetRegions(ctx context.Context) ([]RegionResponse, error) {
	rows, err := s.q.ListRegions(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]RegionResponse, len(rows))
	for i, r := range rows {
		out[i] = RegionResponse{
			Code:                r.Code,
			Name:                r.Name,
			DefaultCurrencyCode: r.DefaultCurrencyCode,
		}
	}
	return out, nil
}
