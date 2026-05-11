package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestTransactionListByUser_OK(t *testing.T) {
	e := newEcho()
	h := NewTransactionHandler(&mockTransactionService{})

	c, rec := newCtx(e, http.MethodGet, "/transactions", "")
	withUser(c, "user-123")
	if err := h.ListByUser(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestTransactionListByUser_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewTransactionHandler(&mockTransactionService{
		listByUserFn: func(_ context.Context, _ string) ([]service.TransactionResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/transactions", "")
	withUser(c, "user-123")
	_ = h.ListByUser(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestTransactionListAll_OK(t *testing.T) {
	e := newEcho()
	h := NewTransactionHandler(&mockTransactionService{})

	c, rec := newCtx(e, http.MethodGet, "/transactions/all", "")
	if err := h.ListAll(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestTransactionListAll_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewTransactionHandler(&mockTransactionService{
		listAllFn: func(_ context.Context) ([]service.TransactionResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/transactions/all", "")
	_ = h.ListAll(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestTransactionPurchase_OK(t *testing.T) {
	e := newEcho()
	h := NewTransactionHandler(&mockTransactionService{})

	c, rec := newCtx(e, http.MethodPost, "/transactions", `{"courseId":"course-1"}`)
	withUser(c, "user-123")
	if err := h.Purchase(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestTransactionPurchase_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewTransactionHandler(&mockTransactionService{})

	c, rec := newCtx(e, http.MethodPost, "/transactions", `{}`)
	withUser(c, "user-123")
	_ = h.Purchase(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestTransactionPurchase_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewTransactionHandler(&mockTransactionService{
		purchaseFn: func(_ context.Context, _, _ string) (*service.TransactionResponse, *service.EnrollmentResponse, error) {
			return nil, nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/transactions", `{"courseId":"course-1"}`)
	withUser(c, "user-123")
	_ = h.Purchase(c)

	assertStatus(t, rec, http.StatusBadRequest)
}
