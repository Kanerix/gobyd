package muex

import (
	"github.com/kanerix/gobyd/pkg/clock"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
	clock clock.VClock
}

// Middleware that initialise our mutual exclusion context.
func CustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Context{c, clock.NewVClock()}
		return next(cc)
	}
}
