package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type UserService struct {
	q *repository.Queries
}

func NewUserService(pool *pgxpool.Pool) *UserService {
	return &UserService{q: repository.New(pool)}
}

func (s *UserService) List(ctx context.Context) ([]UserResponse, error) {
	users, err := s.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]UserResponse, len(users))
	for i, u := range users {
		out[i] = toUserResponse(u, []string{})
	}
	return out, nil
}

func (s *UserService) Delete(ctx context.Context, id string) error {
	uid, err := parseUUID(id)
	if err != nil {
		return err
	}
	return s.q.DeleteUser(ctx, uid)
}
