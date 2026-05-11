package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestCourseList_OK(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	c, rec := newCtx(e, http.MethodGet, "/courses", "")
	if err := h.List(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestCourseList_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		listPublishedFn: func(_ context.Context) ([]service.CourseResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/courses", "")
	_ = h.List(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestCourseListAll_OK(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	c, rec := newCtx(e, http.MethodGet, "/courses/all", "")
	if err := h.ListAll(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestCourseListAll_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		listAllFn: func(_ context.Context) ([]service.CourseResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/courses/all", "")
	_ = h.ListAll(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestCourseGetBySlug_OK(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	c, rec := newCtx(e, http.MethodGet, "/courses/my-course", "")
	c.SetParamNames("slug")
	c.SetParamValues("my-course")
	if err := h.GetBySlug(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestCourseGetBySlug_NotFound(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		getBySlugFn: func(_ context.Context, _ string) (*service.CourseResponse, error) {
			return nil, nil
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/courses/no-such-course", "")
	c.SetParamNames("slug")
	c.SetParamValues("no-such-course")
	_ = h.GetBySlug(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestCourseGetBySlug_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		getBySlugFn: func(_ context.Context, _ string) (*service.CourseResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/courses/my-course", "")
	c.SetParamNames("slug")
	c.SetParamValues("my-course")
	_ = h.GetBySlug(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestCourseUpsert_OK(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	body := `{
		"title":"New Course","slug":"new-course","description":"desc","summary":"sum",
		"category":"Dev","level":"Beginner","instructor":"Alice","instructorBio":"bio",
		"cover":"NC","status":"published","listPrice":49
	}`
	c, rec := newCtx(e, http.MethodPut, "/courses", body)
	if err := h.Upsert(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestCourseUpsert_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	// missing required fields
	c, rec := newCtx(e, http.MethodPut, "/courses", `{"title":"Only Title"}`)
	_ = h.Upsert(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestCourseUpsert_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		upsertFn: func(_ context.Context, _ service.UpsertCourseParams) (*service.CourseResponse, error) {
			return nil, errMock
		},
	})

	body := `{
		"title":"New Course","slug":"new-course","description":"desc","summary":"sum",
		"category":"Dev","level":"Beginner","instructor":"Alice","instructorBio":"bio",
		"cover":"NC","status":"published","listPrice":49
	}`
	c, rec := newCtx(e, http.MethodPut, "/courses", body)
	_ = h.Upsert(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestCourseDelete_OK(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	c, rec := newCtx(e, http.MethodDelete, "/courses/c1", "")
	c.SetParamNames("id")
	c.SetParamValues("c1")
	if err := h.Delete(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusNoContent)
}

func TestCourseDelete_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		deleteFn: func(_ context.Context, _ string) error { return errMock },
	})

	c, rec := newCtx(e, http.MethodDelete, "/courses/c1", "")
	c.SetParamNames("id")
	c.SetParamValues("c1")
	_ = h.Delete(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestCourseSetStatus_OK(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	c, rec := newCtx(e, http.MethodPatch, "/courses/c1/status", `{"status":"published"}`)
	c.SetParamNames("id")
	c.SetParamValues("c1")
	if err := h.SetStatus(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestCourseSetStatus_ValidationError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{})

	c, rec := newCtx(e, http.MethodPatch, "/courses/c1/status", `{}`)
	c.SetParamNames("id")
	c.SetParamValues("c1")
	_ = h.SetStatus(c)

	assertStatus(t, rec, http.StatusBadRequest)
}

func TestCourseSetStatus_NotFound(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		setStatusFn: func(_ context.Context, _, _ string) (*service.CourseResponse, error) {
			return nil, nil
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/courses/c1/status", `{"status":"published"}`)
	c.SetParamNames("id")
	c.SetParamValues("c1")
	_ = h.SetStatus(c)

	assertStatus(t, rec, http.StatusNotFound)
}

func TestCourseSetStatus_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewCourseHandler(&mockCourseService{
		setStatusFn: func(_ context.Context, _, _ string) (*service.CourseResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodPatch, "/courses/c1/status", `{"status":"published"}`)
	c.SetParamNames("id")
	c.SetParamValues("c1")
	_ = h.SetStatus(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
