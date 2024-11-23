package mutex

import "github.com/labstack/echo/v4"

func (h *Handler) Register(g *echo.Group) {
	g.Use(h.MutexContext)

	muex := g.Group("/muex")
	{
		muex.Use(h.MergeClocks)
		muex.Use(h.HandlerEvent)
		muex.GET("/access/request", h.RequestAccess)
		muex.POST("/access/release", h.ReleaseAccess)
	}
}
