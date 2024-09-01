package routes

import (
	handler "monte-carlo-ingestion/internal/handlers"
	"monte-carlo-ingestion/pkg/config"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, cfg *config.Config) {
	h := handler.NewIngestHandler(cfg)
	e.POST("/ingest", h.Ingest)
}
