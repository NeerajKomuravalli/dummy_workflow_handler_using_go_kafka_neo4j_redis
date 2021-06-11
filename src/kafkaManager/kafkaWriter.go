package kafkamanager

import (
	"context"
	"log"
	"os"
	"path/filepath"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/segmentio/kafka-go"
)

type KafkaWriter struct {
	Client *kafka.Writer
}

func GetKafkaWriter() *KafkaWriter {
	if _, err := os.Stat(globalvariables.KafkaLogFolderPath); os.IsNotExist(err) {
		os.MkdirAll(globalvariables.KafkaLogFolderPath, 0777)
	}
	logFile, err := os.OpenFile(
		filepath.Join(globalvariables.KafkaLogFolderPath, globalvariables.KafkaLogFileName),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		globalvariables.KafkaLogFilePermissionCode,
	)
	if err != nil {
		log.Panic(err)
	}
	logger := log.New(logFile, "kafka writer: ", log.Ldate|log.Ltime|log.Lshortfile)
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
