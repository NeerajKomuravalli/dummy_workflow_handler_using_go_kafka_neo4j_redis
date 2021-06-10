package kafkamanager

import (
	"context"
	"log"
	"os"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	Client *kafka.Writer
}

func GetKafkaWriter() *KafkaWriter {
	logger := log.New(os.Stdout, "kafka writer: ", 0)
	kw := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{globalvariables.KafkaBorkerAddress},
		Topic:   globalvariables.KafkaTopicName,
		// assign the logger to the writer
		Logger: logger,
	})
	return &KafkaWriter{kw}
}

func (kw *KafkaWriter) WriteMessages(ctx context.Context, key []byte, value []byte) error {
	err := kw.Client.WriteMessages(ctx, kafka.Message{
		// Key
		Key: key,
		// Payload
		Value: value,
	})
	return err
}
