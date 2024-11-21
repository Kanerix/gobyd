package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Handler for placing a bid.
func (h *Handler) PostBid(c echo.Context) error {
	fmt.Println(c.Request().URL)
	return nil
}

// Handler for getting the current bid.
func (h *Handler) GetBid(c echo.Context) error {
	fmt.Println(c.Request().URL)
	return nil
}

// Handler that gets the result of an auction.
func (h *Handler) GetResult(c echo.Context) error {
	fmt.Println(c.Request().URL)
	return nil
}
