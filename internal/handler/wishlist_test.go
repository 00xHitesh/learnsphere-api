package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestWishlistListByUser_OK(t *testing.T) {
	e := newEcho()
	h := NewWishlistHandler(&mockWishlistService{})

	c, rec := newCtx(e, http.MethodGet, "/wishlist", "")
	withUser(c, "user-123")
	if err := h.ListByUser(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestWishlistListByUser_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewWishlistHandler(&mockWishlistService{
		listByUserFn: func(_ context.Context, _ string) ([]service.WishlistItemResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/wishlist", "")
	withUser(c, "user-123")
	_ = h.ListByUser(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestWishlistToggle_OK(t *testing.T) {
	e := newEcho()
	h := NewWishlistHandler(&mockWishlistService{})

	c, rec := newCtx(e, http.MethodPost, "/wishlist/toggle", `{"courseId":"course-1"}`)
	withUser(c, "user-123")
	if err := h.Toggle(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestWishlistToggle_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewWishlistHandler(&mockWishlistService{})

	c, rec := newCtx(e, http.MethodPost, "/wishlist/toggle", `{}`)
	withUser(c, "user-123")
	_ = h.Toggle(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestWishlistToggle_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewWishlistHandler(&mockWishlistService{
		toggleFn: func(_ context.Context, _, _ string) (bool, error) {
			return false, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/wishlist/toggle", `{"courseId":"course-1"}`)
	withUser(c, "user-123")
	_ = h.Toggle(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestWishlistToggle_RemovesItem(t *testing.T) {
	e := newEcho()
	h := NewWishlistHandler(&mockWishlistService{
		toggleFn: func(_ context.Context, _, _ string) (bool, error) {
			return false, nil
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/wishlist/toggle", `{"courseId":"course-1"}`)
	withUser(c, "user-123")
	if err := h.Toggle(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	body := decodeBody(t, rec)
	data, _ := body["data"].(map[string]interface{})
	if wishlisted, _ := data["wishlisted"].(bool); wishlisted {
		t.Errorf("expected wishlisted=false after removal")
	}
}
