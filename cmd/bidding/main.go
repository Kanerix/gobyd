package main

import (
	"os"

	"github.com/kanerix/gobyd/pkg/gobyd"
	"github.com/kanerix/gobyd/pkg/muex"
	"github.com/kanerix/gobyd/pkg/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(gobyd.CustomContext)
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")

	rest := rest.NewHandler()
	rest.Register(v1)

	muex := muex.NewHandler()
	muex.Register(v1)

	addr := os.Getenv("SERVICE_ADDR")
	if len(addr) < 1 {
		addr = "localhost:8080"
	}

	e.Logger.Fatal(e.Start(addr))
}
