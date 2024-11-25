package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	Bid struct {
		Bid int `json:"bid"`
	}

	BidAnwser struct {
		Accepted bool `json:"accepted"`
	}
)

func (h *RestHandler) PostBid(c echo.Context) error {
	h.NetworkLock()
	defer h.NetworkUnlock()
	return c.JSON(http.StatusOK, BidAnwser{Accepted: true})
}

func (h *RestHandler) GetBid(c echo.Context) error {
	return nil
}

func (h *RestHandler) GetResult(c echo.Context) error {
	return nil
}
