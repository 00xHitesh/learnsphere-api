package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/hitnarkhede/learnsphere-api/pkg/response"
)

type LocaleHandler struct {
	svc localeService
}

func NewLocaleHandler(svc localeService) *LocaleHandler {
	return &LocaleHandler{svc: svc}
}

func (h *LocaleHandler) RegisterRoutes(v1 *echo.Group) {
	g := v1.Group("/locale")
	g.GET("/currencies", h.GetCurrencies)
	g.GET("/regions", h.GetRegions)
}

// GetCurrencies godoc
//
//	@Summary		List currencies
//	@Description	Returns all supported currencies
//	@Tags			Locale
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]service.CurrencyResponse}
//	@Router			/locale/currencies [get]
func (h *LocaleHandler) GetCurrencies(c echo.Context) error {
	data, err := h.svc.GetCurrencies(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, data)
}

// GetRegions godoc
//
//	@Summary		List regions
//	@Description	Returns all supported regions
//	@Tags			Locale
//	@Produce		json
//	@Success		200	{object}	response.Response{data=[]service.RegionResponse}
//	@Router			/locale/regions [get]
func (h *LocaleHandler) GetRegions(c echo.Context) error {
	data, err := h.svc.GetRegions(c.Request().Context())
	if err != nil {
		return response.InternalError(c, "")
	}
	return response.OK(c, data)
}
