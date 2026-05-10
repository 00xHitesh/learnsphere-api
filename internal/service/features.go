package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type FeatureService struct {
	q *repository.Queries
}

func NewFeatureService(pool *pgxpool.Pool) *FeatureService {
	return &FeatureService{q: repository.New(pool)}
}

type FeatureFlagResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	UpdatedAt   string `json:"updatedAt"`
}

func (s *FeatureService) List(ctx context.Context) ([]FeatureFlagResponse, error) {
	rows, err := s.q.ListFeatureFlags(ctx)
	if err != nil {
		return nil, err
	}
	return toFeatureResponses(rows), nil
}

func (s *FeatureService) Toggle(ctx context.Context, name string) ([]FeatureFlagResponse, error) {
	_, err := s.q.ToggleFeatureFlag(ctx, name)
	if err != nil {
		return nil, err
	}
	return s.List(ctx)
}

func toFeatureResponses(rows []repository.FeatureFlag) []FeatureFlagResponse {
	out := make([]FeatureFlagResponse, len(rows))
	for i, f := range rows {
		out[i] = FeatureFlagResponse{
			ID:          uuidToString(f.ID),
			Name:        f.Name,
			Label:       f.Label,
			Description: f.Description,
			Enabled:     f.Enabled,
			UpdatedAt:   f.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		}
	}
	return out
}
