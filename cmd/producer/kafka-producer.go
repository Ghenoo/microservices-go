package producer

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	writer *kafka.Writer
}

func NewKafkaWriter(brokerAddress, topic string) *KafkaWriter {
	return &KafkaWriter{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokerAddress),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (kw *KafkaWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	return kw.writer.WriteMessages(ctx, msgs...)
}

func (kw *KafkaWriter) Close() error {
	return kw.writer.Close()
}

func NewMessage(key, value []byte) kafka.Message {
	return kafka.Message{
		Key:   key,
		Value: value,
	}
}
