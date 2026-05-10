package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type StatsService struct {
	q    *repository.Queries
	pool *pgxpool.Pool
}

func NewStatsService(pool *pgxpool.Pool) *StatsService {
	return &StatsService{q: repository.New(pool), pool: pool}
}

type OverviewResponse struct {
	TotalUsers       int64            `json:"totalUsers"`
	TotalCourses     int64            `json:"totalCourses"`
	TotalEnrollments int64            `json:"totalEnrollments"`
	Revenue          float64          `json:"revenue"`
	TopCourses       []CourseResponse `json:"topCourses"`
}

type PeriodPoint struct {
	Label string  `json:"label"`
	Value float64 `json:"value"`
}

func (s *StatsService) Overview(ctx context.Context) (*OverviewResponse, error) {
	row, err := s.q.GetAdminOverview(ctx)
	if err != nil {
		return nil, err
	}
	top, err := s.q.GetTopCourses(ctx)
	if err != nil {
		return nil, err
	}
	topResp := make([]CourseResponse, len(top))
	for i, c := range top {
		topResp[i] = toCourseResponse(c)
	}
	return &OverviewResponse{
		TotalUsers:       row.TotalUsers,
		TotalCourses:     row.TotalCourses,
		TotalEnrollments: row.TotalEnrollments,
		Revenue:          numericToFloat(row.Revenue),
		TopCourses:       topResp,
	}, nil
}

func (s *StatsService) RevenueByPeriod(ctx context.Context, period string) ([]PeriodPoint, error) {
	trunc, format, days := periodParams(period)
	sql := fmt.Sprintf(`
		SELECT to_char(date_trunc('%s', created_at), '%s') AS label,
		       COALESCE(SUM(amount), 0) AS value
		FROM transactions
		WHERE status = 'success'
		  AND created_at >= NOW() - INTERVAL '%d days'
		GROUP BY date_trunc('%s', created_at)
		ORDER BY date_trunc('%s', created_at)
	`, trunc, format, days, trunc, trunc)
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PeriodPoint
	for rows.Next() {
		var label string
		var val pgtype.Numeric
		if err := rows.Scan(&label, &val); err != nil {
			return nil, err
		}
		out = append(out, PeriodPoint{Label: label, Value: numericToFloat(val)})
	}
	if out == nil {
		out = []PeriodPoint{}
	}
	return out, nil
}

func (s *StatsService) ActiveUsersByPeriod(ctx context.Context, period string) ([]PeriodPoint, error) {
	trunc, format, days := periodParams(period)
	sql := fmt.Sprintf(`
		SELECT to_char(date_trunc('%s', enrolled_at), '%s') AS label,
		       COUNT(DISTINCT user_id) AS value
		FROM enrollments
		WHERE enrolled_at >= NOW() - INTERVAL '%d days'
		GROUP BY date_trunc('%s', enrolled_at)
		ORDER BY date_trunc('%s', enrolled_at)
	`, trunc, format, days, trunc, trunc)
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []PeriodPoint
	for rows.Next() {
		var p PeriodPoint
		var count int64
		if err := rows.Scan(&p.Label, &count); err != nil {
			return nil, err
		}
		p.Value = float64(count)
		out = append(out, p)
	}
	if out == nil {
		out = []PeriodPoint{}
	}
	return out, nil
}

func periodParams(period string) (trunc, format string, days int) {
	switch period {
	case "weekly":
		return "week", "Mon DD", 28
	case "monthly":
		return "month", "Mon", 180
	default:
		return "day", "Mon DD", 7
	}
}
