package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestEnrollmentListByUser_OK(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	c, rec := newCtx(e, http.MethodGet, "/enrollments", "")
	withUser(c, "user-123")
	if err := h.ListByUser(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestEnrollmentListByUser_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		listByUserFn: func(_ context.Context, _ string) ([]service.EnrollmentResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/enrollments", "")
	withUser(c, "user-123")
	_ = h.ListByUser(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestEnrollmentGet_OK(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	c, rec := newCtx(e, http.MethodGet, "/enrollments/course-1", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId")
	c.SetParamValues("course-1")
	if err := h.Get(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestEnrollmentGet_NotFound(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		getFn: func(_ context.Context, _, _ string) (*service.EnrollmentResponse, error) {
			return nil, nil
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/enrollments/course-x", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId")
	c.SetParamValues("course-x")
	_ = h.Get(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestEnrollmentEnroll_OK(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	c, rec := newCtx(e, http.MethodPost, "/enrollments", `{"courseId":"course-1"}`)
	withUser(c, "user-123")
	if err := h.Enroll(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestEnrollmentEnroll_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	c, rec := newCtx(e, http.MethodPost, "/enrollments", `{}`)
	withUser(c, "user-123")
	_ = h.Enroll(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestEnrollmentEnroll_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		enrollFn: func(_ context.Context, _, _ string) (*service.EnrollmentResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPost, "/enrollments", `{"courseId":"course-1"}`)
	withUser(c, "user-123")
	_ = h.Enroll(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestEnrollmentToggleLessonComplete_OK(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/lessons/lesson-1/toggle", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId", "lessonId")
	c.SetParamValues("course-1", "lesson-1")
	if err := h.ToggleLessonComplete(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestEnrollmentToggleLessonComplete_NotFound(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		toggleLessonFn: func(_ context.Context, _, _, _ string) (*service.EnrollmentResponse, error) {
			return nil, nil
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/lessons/lesson-1/toggle", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId", "lessonId")
	c.SetParamValues("course-1", "lesson-1")
	_ = h.ToggleLessonComplete(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestEnrollmentToggleLessonComplete_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		toggleLessonFn: func(_ context.Context, _, _, _ string) (*service.EnrollmentResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/lessons/lesson-1/toggle", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId", "lessonId")
	c.SetParamValues("course-1", "lesson-1")
	_ = h.ToggleLessonComplete(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestEnrollmentRecordTimeSpent_OK(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/time", `{"seconds":120}`)
	withUser(c, "user-123")
	c.SetParamNames("courseId")
	c.SetParamValues("course-1")
	if err := h.RecordTimeSpent(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestEnrollmentRecordTimeSpent_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	// seconds=0 fails min=1
	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/time", `{"seconds":0}`)
	withUser(c, "user-123")
	c.SetParamNames("courseId")
	c.SetParamValues("course-1")
	_ = h.RecordTimeSpent(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestEnrollmentRecordTimeSpent_NotFound(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		recordTimeSpentFn: func(_ context.Context, _, _ string, _ int32) (*service.EnrollmentResponse, error) {
			return nil, nil
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/time", `{"seconds":60}`)
	withUser(c, "user-123")
	c.SetParamNames("courseId")
	c.SetParamValues("course-1")
	_ = h.RecordTimeSpent(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestEnrollmentMarkQuizPassed_OK(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/quiz/quiz-1/pass", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId", "quizId")
	c.SetParamValues("course-1", "quiz-1")
	if err := h.MarkQuizPassed(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestEnrollmentMarkQuizPassed_NotFound(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		markQuizPassedFn: func(_ context.Context, _, _, _ string) (*service.EnrollmentResponse, error) {
			return nil, nil
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/quiz/quiz-1/pass", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId", "quizId")
	c.SetParamValues("course-1", "quiz-1")
	_ = h.MarkQuizPassed(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestEnrollmentMarkQuizPassed_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewEnrollmentHandler(&mockEnrollmentService{
		markQuizPassedFn: func(_ context.Context, _, _, _ string) (*service.EnrollmentResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/enrollments/course-1/quiz/quiz-1/pass", "")
	withUser(c, "user-123")
	c.SetParamNames("courseId", "quizId")
	c.SetParamValues("course-1", "quiz-1")
	_ = h.MarkQuizPassed(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
