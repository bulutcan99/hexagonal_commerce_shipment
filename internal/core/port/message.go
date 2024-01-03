package port

import "context"

type KafkaProducerService interface {
	ProduceMessage(ctx context.Context, topic string, key []byte, value []byte) error
	Close() error
}

type KafkaConsumerService interface {
	ConsumeMessages(ctx context.Context, topic string, handler func(key []byte, value []byte) error) error
	Close() error
}
