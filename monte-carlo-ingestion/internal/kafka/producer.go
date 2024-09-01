package kafka

import (
	"context"
	"monte-carlo-ingestion/pkg/config"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(cfg *config.Config) *Producer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{cfg.KafkaBroker},
		Topic:   cfg.KafkaTopic,
	})

	return &Producer{writer: writer}
}

func (p *Producer) Publish(ctx context.Context, message []byte) error {
	return p.writer.WriteMessages(ctx, kafka.Message{
		Value: message,
	})
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
