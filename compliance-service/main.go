package main

import (
	handler "compliance-service/internal/handlers"
	"compliance-service/pkg/config"
)

func main() {
	cfg := &config.Config{
		KafkaBroker: "kafka:9092",
		KafkaTopic:  "pii-gdpr-annotations",
	}

	h := handler.NewComplianceHandler(cfg)
	h.Start()
}
