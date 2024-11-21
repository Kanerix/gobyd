package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(g *echo.Group) {
	g.POST("/auction/bid", h.PostBid)
	g.GET("/auction/bid", h.GetBid)
	g.GET("/auction/result", h.GetResult)
}
