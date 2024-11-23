package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	PostBidRequest struct {
		Bid int `json:"bid"`
	}

	PostBidResponse struct {
		Accepted bool `json:"accepted"`
	}

	GetBidResponse struct {
		Bid int `json:"bid"`
	}

	GetResultResponse struct {
		Bid int `json:"bid"`
	}
)

func (h *Handler) PostBid(c echo.Context) error {
	h.NetworkLock()
	defer h.NetworkUnlock()
	return c.JSON(http.StatusOK, PostBidResponse{Accepted: true})
}

func (h *Handler) GetBid(c echo.Context) error {
	return nil
}

func (h *Handler) GetResult(c echo.Context) error {
	return nil
}
