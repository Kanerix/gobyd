package rest

import (
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context
}

// Middleware that initialise our custom context.
func CustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &Context{c}
		return next(cc)
	}
}
