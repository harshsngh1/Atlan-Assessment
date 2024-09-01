package handler

import (
	"context"
	"log"

	"compliance-service/pkg/config"

	"github.com/segmentio/kafka-go"
)

type ComplianceHandler struct {
	reader *kafka.Reader
}

func NewComplianceHandler(cfg *config.Config) *ComplianceHandler {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   cfg.KafkaTopic,
		GroupID: "compliance-consumers",
	})

	return &ComplianceHandler{reader: reader}
}

func (h *ComplianceHandler) Start() {
	for {
		m, err := h.reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		h.notifyExternalSystems(m.Value)
	}
}

func (h *ComplianceHandler) notifyExternalSystems(data []byte) {
	log.Printf("Notifying external systems with: %s", string(data))
}
