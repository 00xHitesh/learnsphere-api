package handler

import (
	"context"
	"net/http"
	"testing"

	"github.com/hitnarkhede/learnsphere-api/internal/service"
)

func TestLocaleGetCurrencies_OK(t *testing.T) {
	e := newEcho()
	h := NewLocaleHandler(&mockLocaleService{})

	c, rec := newCtx(e, http.MethodGet, "/locale/currencies", "")
	if err := h.GetCurrencies(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestLocaleGetCurrencies_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewLocaleHandler(&mockLocaleService{
		getCurrenciesFn: func(_ context.Context) ([]service.CurrencyResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/locale/currencies", "")
	_ = h.GetCurrencies(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}

func TestLocaleGetRegions_OK(t *testing.T) {
	e := newEcho()
	h := NewLocaleHandler(&mockLocaleService{})

	c, rec := newCtx(e, http.MethodGet, "/locale/regions", "")
	if err := h.GetRegions(c); err != nil {
		t.Fatal(err)
	}

	assertStatus(t, rec, http.StatusOK)
	assertSuccess(t, rec)
}

func TestLocaleGetRegions_ServiceError(t *testing.T) {
	e := newEcho()
	h := NewLocaleHandler(&mockLocaleService{
		getRegionsFn: func(_ context.Context) ([]service.RegionResponse, error) {
			return nil, errMock
		},
	})

	c, rec := newCtx(e, http.MethodGet, "/locale/regions", "")
	_ = h.GetRegions(c)

	assertStatus(t, rec, http.StatusInternalServerError)
}
