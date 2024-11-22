package rest

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) ReleaseAccess(c echo.Context) error {
	// cc := c.(*Context)
	return nil
}

func (h *Handler) RequestAccess(c echo.Context) error {
	// cc := c.(*Context)
	return nil
}

func (h *Handler) AskPeer() {
	h.RLock()
	defer h.RUnlock()

	h.clock.TickProcess(h.nodeID)
}
