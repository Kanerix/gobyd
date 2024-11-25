package rest

import "github.com/labstack/echo/v4"

func (h *RestHandler) Register(e *echo.Echo) {
	h.Register(e)

	g := e.Group("/api")
	g.POST("/auction/bid", h.PostBid)
	g.GET("/auction/bid", h.GetBid)
	g.GET("/auction/result", h.GetResult)
}
