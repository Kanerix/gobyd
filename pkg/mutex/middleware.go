package mutex

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/kanerix/gobyd/pkg/clock"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	requesterID    uuid.UUID
	requesterClock clock.VClock
}

func (h *Handler) MutexContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqHeader := c.Request().Header

		requesterID, err := GetNodeID(reqHeader)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error(), err)
		}

		requesterClock, err := clock.FromHeader(reqHeader)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error(), err)
		}

		cc := &Context{c, requesterID, requesterClock}
		res := next(cc)

		h.RLock()
		vcHeader := h.Clock.IntoHeader()
		h.RUnlock()
		resHeader := c.Response().Header()
		for key, values := range vcHeader {
			for _, value := range values {
				resHeader.Add(key, value)
			}
		}

		return res
	}
}

func (h *Handler) MergeClocks(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*Context)

		h.Lock()
		h.Clock.Merge(cc.requesterClock)
		h.Unlock()

		return next(cc)
	}
}

func (h *Handler) HandlerEvent(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.TickClock()
		return next(c)
	}
}

func GetNodeID(h http.Header) (requesterID uuid.UUID, err error) {
	header := h["Requester-ID"]
	for _, value := range header {
		requesterID, err = uuid.Parse(value)
		if err != nil {
			return uuid.UUID{}, ErrHeaderParseFormat
		}
		return requesterID, nil
	}
	return uuid.UUID{}, ErrHeaderMissing
}

var (
	ErrHeaderMissing     = errors.New("missing \"Requester-ID\" header")
	ErrHeaderParseFormat = errors.New("error parsing \"Requester-ID\" header as UUID")
)
