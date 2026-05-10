package service

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type ChatService struct {
	q *repository.Queries
}

func NewChatService(pool *pgxpool.Pool) *ChatService {
	return &ChatService{q: repository.New(pool)}
}

type ChatMessageResponse struct {
	ID        string `json:"id"`
	Role      string `json:"role"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

func (s *ChatService) List(ctx context.Context, userID string, courseID *string) ([]ChatMessageResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	cid := pgtype.UUID{}
	if courseID != nil {
		cid, err = parseUUID(*courseID)
		if err != nil {
			return nil, err
		}
	}
	session, err := s.q.GetOrCreateChatSession(ctx, repository.GetOrCreateChatSessionParams{
		UserID: uid, CourseID: cid,
	})
	if err != nil {
		return nil, err
	}
	msgs, err := s.q.ListChatMessages(ctx, session.ID)
	if err != nil {
		return nil, err
	}
	out := make([]ChatMessageResponse, len(msgs))
	for i, m := range msgs {
		out[i] = ChatMessageResponse{
			ID:        uuidToString(m.ID),
			Role:      string(m.Role),
			Content:   m.Content,
			CreatedAt: m.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		}
	}
	return out, nil
}

func (s *ChatService) Append(ctx context.Context, userID string, courseID *string, role, content string) (*ChatMessageResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	cid := pgtype.UUID{}
	if courseID != nil {
		cid, err = parseUUID(*courseID)
		if err != nil {
			return nil, err
		}
	}
	session, err := s.q.GetOrCreateChatSession(ctx, repository.GetOrCreateChatSessionParams{
		UserID: uid, CourseID: cid,
	})
	if err != nil {
		return nil, err
	}
	msg, err := s.q.InsertChatMessage(ctx, repository.InsertChatMessageParams{
		SessionID: session.ID,
		Role:      repository.ChatRole(role),
		Content:   content,
	})
	if err != nil {
		return nil, err
	}
	resp := ChatMessageResponse{
		ID:        uuidToString(msg.ID),
		Role:      string(msg.Role),
		Content:   msg.Content,
		CreatedAt: msg.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
	return &resp, nil
}

func (s *ChatService) Reset(ctx context.Context, userID string, courseID *string) error {
	uid, err := parseUUID(userID)
	if err != nil {
		return err
	}
	cid := pgtype.UUID{}
	if courseID != nil {
		cid, err = parseUUID(*courseID)
		if err != nil {
			return err
		}
	}
	session, err := s.q.GetOrCreateChatSession(ctx, repository.GetOrCreateChatSessionParams{
		UserID: uid, CourseID: cid,
	})
	if err != nil {
		return err
	}
	return s.q.DeleteChatMessages(ctx, session.ID)
}
