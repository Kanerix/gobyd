package mutex

import "github.com/labstack/echo/v4"

func (h *MutexHandler) Register(e *echo.Echo) {
	e.Use(h.MutexContext)

	muex := e.Group("/muex")
	{
		muex.Use(h.MergeClocks)
		muex.Use(h.HandlerEvent)
		muex.GET("/access/request", h.RequestAccess)
		muex.POST("/access/release", h.ReleaseAccess)
	}
}
