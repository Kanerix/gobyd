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

	name := GetenvDefault("SERVICE_NAME", "bidding")
	addr := GetenvDefault("SERVICE_ADDR", "localhost:8080")
	servers := strings.Split(serversRaw, "\n")

	logging(e, name)

	if handler := rest.NewRestHandler(servers); handler != nil {
		handler.Register(e)
	}

	e.Logger.Fatal(e.Start(addr))
}

func logging(e *echo.Echo, name string) {
	out := os.Stdout
	if file := os.Getenv("SERVICE_LOG_TO_FILE"); len(file) > 0 {
		path := path.Join("logs", fmt.Sprintf("%s.logs", name))
		perms := os.O_RDWR | os.O_CREATE | os.O_APPEND
		f, err := os.OpenFile(path, perms, 0666)
		if err != nil {
			panic(fmt.Sprintf("error opening file: %v", err))
		}
		out = f
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		Output: out,
	}))
}

func GetenvDefault(env string, def string) string {
	v, ok := os.LookupEnv(env)
	if !ok {
		return def
	}
	return v
}
