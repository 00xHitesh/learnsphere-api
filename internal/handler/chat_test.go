package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestChatList_OK(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{})

	c, rec := newCtx(e, http.MethodGet, "/chat", "")
	withUser(c, "user-123")
	if err := h.List(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestChatList_WithCourseID(t *testing.T) {
	e := newEcho()
	var capturedCourseID *string
	h := NewChatHandler(&mockChatService{
		listFn: func(_ context.Context, _ string, courseID *string) ([]service.ChatMessageResponse, error) {
			capturedCourseID = courseID
			return []service.ChatMessageResponse{}, nil
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/chat?courseId=course-1", "")
	withUser(c, "user-123")
	if err := h.List(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	if capturedCourseID == nil || *capturedCourseID != "course-1" {
		t.Errorf("expected courseID=course-1, got %v", capturedCourseID)
	}
}

func TestChatList_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{
		listFn: func(_ context.Context, _ string, _ *string) ([]service.ChatMessageResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/chat", "")
	withUser(c, "user-123")
	_ = h.List(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestChatAppend_OK(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{})

	c, rec := newCtx(e, http.MethodPost, "/chat/messages", `{"role":"user","content":"What is Go?"}`)
	withUser(c, "user-123")
	if err := h.Append(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestChatAppend_ValidationError_InvalidRole(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{})

	// role must be "user" or "assistant"
	c, rec := newCtx(e, http.MethodPost, "/chat/messages", `{"role":"admin","content":"hi"}`)
	withUser(c, "user-123")
	_ = h.Append(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestChatAppend_ValidationError_MissingContent(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{})

	c, rec := newCtx(e, http.MethodPost, "/chat/messages", `{"role":"user"}`)
	withUser(c, "user-123")
	_ = h.Append(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestChatAppend_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{
		appendFn: func(_ context.Context, _ string, _ *string, _, _ string) (*service.ChatMessageResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/chat/messages", `{"role":"user","content":"hello"}`)
	withUser(c, "user-123")
	_ = h.Append(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestChatReset_OK(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{})

	c, rec := newCtx(e, http.MethodDelete, "/chat", "")
	withUser(c, "user-123")
	if err := h.Reset(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusNoContent)
}

func TestChatReset_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewChatHandler(&mockChatService{
		resetFn: func(_ context.Context, _ string, _ *string) error {
			return errMock
		},
	})

	c, rec := newCtx(e, http.MethodDelete, "/chat", "")
	withUser(c, "user-123")
	_ = h.Reset(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
