package service

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type LiveSessionService struct {
	q *repository.Queries
}

func NewLiveSessionService(pool *pgxpool.Pool) *LiveSessionService {
	return &LiveSessionService{q: repository.New(pool)}
}

type LiveSessionResponse struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Instructor      string   `json:"instructor"`
	Description     string   `json:"description"`
	ScheduledAt     string   `json:"scheduledAt"`
	DurationMinutes int32    `json:"durationMinutes"`
	Topic           string   `json:"topic"`
	JoinURL         string   `json:"joinUrl"`
	CourseID        *string  `json:"courseId"`
	RegisteredCount int32    `json:"registeredCount"`
	MaxAttendees    int32    `json:"maxAttendees"`
	ListPrice       *float64 `json:"listPrice"`
	SalePrice       *float64 `json:"salePrice"`
	Status          string   `json:"status"`
	CreatedAt       string   `json:"createdAt"`
	UpdatedAt       string   `json:"updatedAt"`
}

type UpsertLiveSessionParams struct {
	ID              string
	Title           string
	Instructor      string
	Description     string
	ScheduledAt     pgtype.Timestamptz
	DurationMinutes int32
	Topic           string
	JoinURL         string
	CourseID        *string
	RegisteredCount int32
	MaxAttendees    int32
	ListPrice       *float64
	SalePrice       *float64
	Status          string
}

func (s *LiveSessionService) ListPublished(ctx context.Context) ([]LiveSessionResponse, error) {
	rows, err := s.q.ListPublishedLiveSessions(ctx)
	if err != nil {
		return nil, err
	}
	return toLiveSessionResponses(rows), nil
}

func (s *LiveSessionService) ListAll(ctx context.Context) ([]LiveSessionResponse, error) {
	rows, err := s.q.ListAllLiveSessions(ctx)
	if err != nil {
		return nil, err
	}
	return toLiveSessionResponses(rows), nil
}

func (s *LiveSessionService) Upsert(ctx context.Context, p UpsertLiveSessionParams) (*LiveSessionResponse, error) {
	cid := pgtype.UUID{}
	if p.CourseID != nil {
		var err error
		cid, err = parseUUID(*p.CourseID)
		if err != nil {
			return nil, err
		}
	}
	listPrice := optFloatToNumeric(p.ListPrice)
	salePrice := optFloatToNumeric(p.SalePrice)
	status := repository.ContentStatus(p.Status)

	var row repository.LiveSession
	var err error
	if p.ID == "" {
		row, err = s.q.InsertLiveSession(ctx, repository.InsertLiveSessionParams{
			Title: p.Title, Instructor: p.Instructor, Description: p.Description,
			ScheduledAt: p.ScheduledAt, DurationMinutes: p.DurationMinutes,
			Topic: p.Topic, JoinUrl: p.JoinURL, CourseID: cid,
			RegisteredCount: p.RegisteredCount, MaxAttendees: p.MaxAttendees,
			ListPrice: listPrice, SalePrice: salePrice, Status: status,
		})
	} else {
		uid, parseErr := parseUUID(p.ID)
		if parseErr != nil {
			return nil, parseErr
		}
		row, err = s.q.UpdateLiveSession(ctx, repository.UpdateLiveSessionParams{
			ID: uid, Title: p.Title, Instructor: p.Instructor, Description: p.Description,
			ScheduledAt: p.ScheduledAt, DurationMinutes: p.DurationMinutes,
			Topic: p.Topic, JoinUrl: p.JoinURL, CourseID: cid,
			RegisteredCount: p.RegisteredCount, MaxAttendees: p.MaxAttendees,
			ListPrice: listPrice, SalePrice: salePrice, Status: status,
		})
		if err != nil && errors.Is(err, pgx.ErrNoRows) {
			row, err = s.q.InsertLiveSession(ctx, repository.InsertLiveSessionParams{
				Title: p.Title, Instructor: p.Instructor, Description: p.Description,
				ScheduledAt: p.ScheduledAt, DurationMinutes: p.DurationMinutes,
				Topic: p.Topic, JoinUrl: p.JoinURL, CourseID: cid,
				RegisteredCount: p.RegisteredCount, MaxAttendees: p.MaxAttendees,
				ListPrice: listPrice, SalePrice: salePrice, Status: status,
			})
		}
	}
	if err != nil {
		return nil, err
	}
	resp := toLiveSessionResponse(row)
	return &resp, nil
}

func (s *LiveSessionService) Delete(ctx context.Context, id string) error {
	uid, err := parseUUID(id)
	if err != nil {
		return err
	}
	return s.q.DeleteLiveSession(ctx, uid)
}

func (s *LiveSessionService) SetStatus(ctx context.Context, id, status string) (*LiveSessionResponse, error) {
	uid, err := parseUUID(id)
	if err != nil {
		return nil, err
	}
	row, err := s.q.UpdateLiveSessionStatus(ctx, repository.UpdateLiveSessionStatusParams{
		ID:     uid,
		Status: repository.ContentStatus(status),
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	resp := toLiveSessionResponse(row)
	return &resp, nil
}

func toLiveSessionResponses(rows []repository.LiveSession) []LiveSessionResponse {
	out := make([]LiveSessionResponse, len(rows))
	for i, r := range rows {
		out[i] = toLiveSessionResponse(r)
	}
	return out
}

func toLiveSessionResponse(s repository.LiveSession) LiveSessionResponse {
	var cid *string
	if s.CourseID.Valid {
		str := uuidToString(s.CourseID)
		cid = &str
	}
	return LiveSessionResponse{
		ID:              uuidToString(s.ID),
		Title:           s.Title,
		Instructor:      s.Instructor,
		Description:     s.Description,
		ScheduledAt:     s.ScheduledAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		DurationMinutes: s.DurationMinutes,
		Topic:           s.Topic,
		JoinURL:         s.JoinUrl,
		CourseID:        cid,
		RegisteredCount: s.RegisteredCount,
		MaxAttendees:    s.MaxAttendees,
		ListPrice:       nullNumericToPtr(s.ListPrice),
		SalePrice:       nullNumericToPtr(s.SalePrice),
		Status:          string(s.Status),
		CreatedAt:       s.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:       s.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
}
