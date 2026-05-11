package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/config"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
	"github.com/hitnarkhede/learnsphere-api/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	q   *repository.Queries
	cfg *config.Config
}

func NewAuthService(pool *pgxpool.Pool, cfg *config.Config) *AuthService {
	return &AuthService{q: repository.New(pool), cfg: cfg}
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserResponse struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	Role         string   `json:"role"`
	AvatarURL    *string  `json:"avatarUrl"`
	Bio          *string  `json:"bio"`
	Points       int32    `json:"points"`
	RegionCode   string   `json:"regionCode"`
	CurrencyCode string   `json:"currencyCode"`
	Interests    []string `json:"interests"`
	CreatedAt    string   `json:"createdAt"`
}

func (s *AuthService) Login(ctx context.Context, email, password string) (*UserResponse, *TokenPair, error) {
	user, err := s.q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, fmt.Errorf("invalid email or password")
		}
		return nil, nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, nil, fmt.Errorf("invalid email or password")
	}
	return s.buildAuthResponse(ctx, user)
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) (*UserResponse, *TokenPair, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, nil, err
	}
	user, err := s.q.CreateUser(ctx, repository.CreateUserParams{
		Name:         name,
		Lower:        email,
		PasswordHash: string(hash),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("an account with that email already exists")
	}
	return s.buildAuthResponse(ctx, user)
}

func (s *AuthService) GetUser(ctx context.Context, userID string) (*UserResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	user, err := s.q.GetUserByID(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	interests, _ := s.q.GetUserInterests(ctx, user.ID)
	resp := toUserResponse(user, interests)
	return &resp, nil
}

func (s *AuthService) UpdateInterests(ctx context.Context, userID string, interests []string) (*UserResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	if err := s.q.DeleteUserInterests(ctx, uid); err != nil {
		return nil, err
	}
	for _, cat := range interests {
		_ = s.q.InsertUserInterest(ctx, repository.InsertUserInterestParams{UserID: uid, Category: cat})
	}
	user, err := s.q.GetUserByID(ctx, uid)
	if err != nil {
		return nil, err
	}
	resp := toUserResponse(user, interests)
	return &resp, nil
}

func (s *AuthService) UpdateLocale(ctx context.Context, userID, regionCode, currencyCode string) (*UserResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	user, err := s.q.UpdateUserLocale(ctx, repository.UpdateUserLocaleParams{
		ID:           uid,
		RegionCode:   regionCode,
		CurrencyCode: currencyCode,
	})
	if err != nil {
		return nil, err
	}
	interests, _ := s.q.GetUserInterests(ctx, uid)
	resp := toUserResponse(user, interests)
	return &resp, nil
}

func (s *AuthService) RequestReset(ctx context.Context, email string) (string, error) {
	user, err := s.q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("no account with that email")
		}
		return "", err
	}
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	token := hex.EncodeToString(b)
	_, err = s.q.CreateResetToken(ctx, repository.CreateResetTokenParams{
		UserID: user.ID,
		Token:  token,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, token, password string) error {
	rt, err := s.q.GetActiveResetToken(ctx, token)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fmt.Errorf("invalid or expired reset link")
		}
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err := s.q.UpdateUserPassword(ctx, repository.UpdateUserPasswordParams{
		ID:           rt.UserID,
		PasswordHash: string(hash),
	}); err != nil {
		return err
	}
	return s.q.MarkResetTokenUsed(ctx, rt.ID)
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error) {
	claims, err := jwt.ValidatToken(refreshToken, s.cfg.JWTRefreshSecret)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired refresh token")
	}
	uid, err := parseUUID(claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("invalid token claims")
	}
	user, err := s.q.GetUserByID(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	pair, err := jwt.GenerateTokenPair(
		uuidToString(user.ID), user.Email, string(user.Role),
		s.cfg.JWTAccessSecret, s.cfg.JWTRefreshSecret,
		s.cfg.JWTAccessExpMin, s.cfg.JWTRefreshExpDay,
	)
	if err != nil {
		return nil, err
	}
	return &TokenPair{AccessToken: pair.AccessToken, RefreshToken: pair.RefreshToken}, nil
}

func (s *AuthService) buildAuthResponse(ctx context.Context, user repository.User) (*UserResponse, *TokenPair, error) {
	interests, _ := s.q.GetUserInterests(ctx, user.ID)
	pair, err := jwt.GenerateTokenPair(
		uuidToString(user.ID), user.Email, string(user.Role),
		s.cfg.JWTAccessSecret, s.cfg.JWTRefreshSecret,
		s.cfg.JWTAccessExpMin, s.cfg.JWTRefreshExpDay,
	)
	if err != nil {
		return nil, nil, err
	}
	resp := toUserResponse(user, interests)
	return &resp, &TokenPair{AccessToken: pair.AccessToken, RefreshToken: pair.RefreshToken}, nil
}