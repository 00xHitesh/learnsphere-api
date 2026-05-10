package service

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/hitnarkhede/learnsphere-api/internal/repository"
)

type TransactionService struct {
	q *repository.Queries
}

func NewTransactionService(pool *pgxpool.Pool) *TransactionService {
	return &TransactionService{q: repository.New(pool)}
}

type TransactionResponse struct {
	ID        string  `json:"id"`
	UserID    *string `json:"userId"`
	CourseID  *string `json:"courseId"`
	Amount    float64 `json:"amount"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"createdAt"`
}

func (s *TransactionService) ListByUser(ctx context.Context, userID string) ([]TransactionResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, err
	}
	rows, err := s.q.ListTransactionsByUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return toTransactionResponses(rows), nil
}

func (s *TransactionService) ListAll(ctx context.Context) ([]TransactionResponse, error) {
	rows, err := s.q.ListAllTransactions(ctx)
	if err != nil {
		return nil, err
	}
	return toTransactionResponses(rows), nil
}

func (s *TransactionService) Purchase(ctx context.Context, userID, courseID string) (*TransactionResponse, *EnrollmentResponse, error) {
	uid, err := parseUUID(userID)
	if err != nil {
		return nil, nil, err
	}
	cid, err := parseUUID(courseID)
	if err != nil {
		return nil, nil, err
	}
	course, err := s.q.GetCourseByID(ctx, cid)
	if err != nil {
		return nil, nil, fmt.Errorf("course not found")
	}

	amount := course.ListPrice
	if course.SalePrice.Valid {
		amount = course.SalePrice // use discounted price when available
	}

	tx, err := s.q.InsertTransaction(ctx, repository.InsertTransactionParams{
		UserID:   uid,
		CourseID: cid,
		Amount:   amount,
		Status:   repository.TransactionStatusSuccess,
	})
	if err != nil {
		return nil, nil, err
	}

	enrollment, err := s.q.InsertEnrollment(ctx, repository.InsertEnrollmentParams{
		UserID: uid, CourseID: cid,
	})
	if err != nil {
		return nil, nil, err
	}

	txResp := toTransactionResponse(tx)
	eResp := toEnrollmentResponse(enrollment, []string{})
	return &txResp, &eResp, nil
}

func toTransactionResponses(rows []repository.Transaction) []TransactionResponse {
	out := make([]TransactionResponse, len(rows))
	for i, r := range rows {
		out[i] = toTransactionResponse(r)
	}
	return out
}

func toTransactionResponse(t repository.Transaction) TransactionResponse {
	var uid, cid *string
	if t.UserID.Valid {
		s := uuidToString(t.UserID)
		uid = &s
	}
	if t.CourseID.Valid {
		s := uuidToString(t.CourseID)
		cid = &s
	}
	return TransactionResponse{
		ID:        uuidToString(t.ID),
		UserID:    uid,
		CourseID:  cid,
		Amount:    numericToFloat(t.Amount),
		Status:    string(t.Status),
		CreatedAt: t.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
}
