package main

import (
	"monte-carlo-ingestion/internal/routes"
	"monte-carlo-ingestion/pkg/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	cfg := config.LoadConfig()

	routes.RegisterRoutes(e, cfg)

	e.Logger.Fatal(e.Start(":8080"))
}
