package service

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

func parseUUID(s string) (pgtype.UUID, error) {
	var uid pgtype.UUID
	if err := uid.Scan(s); err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid id: %w", err)
	}
	return uid, nil
}

func uuidToString(u pgtype.UUID) string {
	if !u.Valid {
		return ""
	}
	b := u.Bytes
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func nullTextToPtr(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}

func nullNumericToPtr(n pgtype.Numeric) *float64 {
	if !n.Valid {
		return nil
	}
	f, err := n.Float64Value()
	if err != nil || !f.Valid {
		return nil
	}
	return &f.Float64
}

func numericToFloat(n pgtype.Numeric) float64 {
	f, err := n.Float64Value()
	if err != nil || !f.Valid {
		return 0
	}
	return f.Float64
}

func toUserResponse(user repository.User, interests []string) UserResponse {
	if interests == nil {
		interests = []string{}
	}
	return UserResponse{
		ID:           uuidToString(user.ID),
		Name:         user.Name,
		Email:        user.Email,
		Role:         string(user.Role),
		AvatarURL:    nullTextToPtr(user.AvatarUrl),
		Bio:          nullTextToPtr(user.Bio),
		Points:       user.Points,
		RegionCode:   user.RegionCode,
		CurrencyCode: user.CurrencyCode,
		Interests:    interests,
		CreatedAt:    user.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
}
