package kafkamanager

import (
	"context"
	"log"
	"os"
	"path/filepath"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/segmentio/kafka-go"
)

type KafkaReader struct {
	Client *kafka.Reader
}

func GetKafkaReader() *KafkaReader {
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
	logger := log.New(logFile, "kafka reader: ", log.Ldate|log.Ltime|log.Lshortfile)
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
