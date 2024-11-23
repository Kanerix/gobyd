package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	_ "embed"

	"github.com/kanerix/gobyd/pkg/rest"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed servers.txt
var serversRaw string

func main() {
	e := echo.New()

	name := os.Getenv("SERVICE_NAME")
	if len(name) < 1 {
		name = "bidding"
	}

	logging(e, name)

	servers := strings.Split(serversRaw, "\n")

	api := e.Group("/api")
	rest := rest.NewHandler(servers)
	rest.Register(api)

	addr := os.Getenv("SERVICE_ADDR")
	if len(addr) < 1 {
		addr = "localhost:8080"
	}

	e.Logger.Fatal(e.Start(addr))
}

func logging(e *echo.Echo, name string) {
	out := os.Stdout
	if file := os.Getenv("SERVICE_LOG_TO_FILE"); len(file) > 0 {
		f, err := os.OpenFile(path.Join("logs", name+".logs"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(fmt.Sprintf("error opening file: %v", err))
		}
		defer f.Close()
		out = f
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: out,
	}))
}
