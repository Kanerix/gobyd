package rest

import "github.com/labstack/echo/v4"

func (h *Handler) Register(g *echo.Group) {
	g.Use(h.MuexContext)
	g.Use(h.MuexTickClock)

	g.POST("/auction/bid", h.PostBid)
	g.GET("/auction/bid", h.GetBid)
	g.GET("/auction/result", h.GetResult)

	muex := g.Group("/muex")
	{
		muex.POST("/access/release", h.ReleaseAccess)
		muex.GET("/access/request", h.RequestAccess)
	}
}
