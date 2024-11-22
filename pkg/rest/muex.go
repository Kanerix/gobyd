package rest

import (
	"net/http"

	"github.com/kanerix/gobyd/pkg/clock"
	"github.com/labstack/echo/v4"
)

type MuexContext struct {
	echo.Context
	clock clock.VClock
}

// Middleware that initialise the mutual exclusion context.
func (h *Handler) MuexContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqHeader := c.Request().Header
		vc, err := clock.FromHeader(reqHeader)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		cc := &MuexContext{c, vc}
		res := next(cc)

		vcHeader := vc.IntoHeader()
		resHeader := cc.Response().Header()
		for key, values := range vcHeader {
			for _, value := range values {
				resHeader.Add(key, value)
			}
		}

		return res
	}
}

// Middleware that ticks the local vector clock.
func (h *Handler) MuexTickClock(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*MuexContext)
		cc.clock.TickProcess(h.nodeID)
		return next(c)
	}
}
