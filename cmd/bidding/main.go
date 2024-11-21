package main

import (
	"os"

	"github.com/kanerix/gobyd/handler"
	"github.com/kanerix/gobyd/pkg/gobyd"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(gobyd.CustomContext)
	e.Use(middleware.Logger())

	v1 := e.Group("/api/v1")
	handler := handler.NewHandler()
	handler.Register(v1)

	muex := muex.NewHandler()
	muex.Register(v1)

	addr := os.Getenv("SERVICE_ADDR")
	if len(addr) < 1 {
		addr = "localhost:8080"
	}

	e.Logger.Fatal(e.Start(addr))
}
