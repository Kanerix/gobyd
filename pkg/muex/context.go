package muex

import (
	"net/http"

	"github.com/kanerix/gobyd/pkg/clock"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	clock clock.VClock
}

// Middleware that initialise the mutual exclusion context.
func (h *Handler) MuexContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		vc, err := clock.FromHeader(c.Request().Header)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		vc.TickProcess(h.nodeID)

		cc := &Context{c, vc}
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
