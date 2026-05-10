package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type WishlistService struct {
	q *repository.Queries
}

func NewWishlistService(pool *pgxpool.Pool) *WishlistService {
	return &WishlistService{q: repository.New(pool)}
}

type WishlistItemResponse struct {
	UserID    string `json:"userId"`
	CourseID  string `json:"courseId"`
	AddedAt   string `json:"addedAt"`
}

func (s *WishlistService) ListByUser(ctx context.Context, userID string) ([]WishlistItemResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	rows, err := s.q.ListWishlistByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	out := make([]WishlistItemResponse, len(rows))
	for i, r := range rows {
		out[i] = WishlistItemResponse{
			UserID:   uuidToString(r.UserID),
			CourseID: uuidToString(r.CourseID),
			AddedAt:  r.AddedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		}
	}
	return out, nil
}

func (s *WishlistService) Toggle(ctx context.Context, userID, courseID string) (bool, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return false, err
	}
	cid, err := parseUUID(courseID)
	if err != nil {
		return false, err
	}
	_, err = s.q.GetWishlistItem(ctx, repository.GetWishlistItemParams{UserID: uid, CourseID: cid})
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return false, err
	}
	if err == nil {
		_ = s.q.DeleteWishlistItem(ctx, repository.DeleteWishlistItemParams{UserID: uid, CourseID: cid})
		return false, nil
	}
	_ = s.q.InsertWishlistItem(ctx, repository.InsertWishlistItemParams{UserID: uid, CourseID: cid})
	return true, nil
}
