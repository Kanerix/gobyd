package muex

import "github.com/labstack/echo/v4"

func (h *Handler) Register(e *echo.Echo) {
	e.Use(h.MuexContext)
	e.POST("/access/release", h.ReleaseAccess)
	e.GET("/access/request", h.RequestAccess)
}
