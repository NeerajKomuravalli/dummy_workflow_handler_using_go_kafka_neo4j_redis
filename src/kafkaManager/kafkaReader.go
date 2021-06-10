package kafkamanager

import (
	"context"
	"log"
	"os"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	Client *kafka.Reader
}

func GetKafkaReader() *KafkaReader {
	logger := log.New(os.Stdout, "kafka reader: ", 0)
	kr := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{globalvariables.KafkaBorkerAddress},
		Topic:   globalvariables.KafkaTopicName,
		GroupID: globalvariables.KafkaGroupId,
		// assign the logger to the reader
		Logger: logger,
	})
	return &KafkaReader{kr}
}

func (kr *KafkaReader) ReadMessages(ctx context.Context) ([]byte, error) {
	msg, err := kr.Client.ReadMessage(ctx)
	return msg.Value, err
}
