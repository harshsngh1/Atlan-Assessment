package handler

import (
	"context"
	"encoding/json"
	"monte-carlo-ingestion/internal/kafka"
	"monte-carlo-ingestion/pkg/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IngestHandler struct {
	cfg    *config.Config
	writer *kafka.Producer
}

func NewIngestHandler(cfg *config.Config) *IngestHandler {
	writer := kafka.NewProducer(cfg)
	return &IngestHandler{
		cfg:    cfg,
		writer: writer,
	}
}

func (h *IngestHandler) Ingest(c echo.Context) error {
	var data map[string]interface{}
	if err := c.Bind(&data); err != nil {
		c.Logger().Error("Failed to bind request data: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	message, err := json.Marshal(data)
	if err != nil {
		c.Logger().Error("Failed to marshal JSON: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to process data")
	}

	err = h.writer.Publish(context.Background(), message)
	if err != nil {
		c.Logger().Error("Failed to publish message to Kafka: ", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to send data to Kafka")
	}

	return c.JSON(http.StatusOK, "Received")
}
