package main

// - Take Device data from Kafka
// - Trigger relevant workflows (as of now dummy workflow), if any
// - update neo4j

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	kafkamanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/kafkaManager"
	log "github.com/sirupsen/logrus"
)

func setUpLogger() {
	if _, err := os.Stat(globalvariables.WorkflowLogFolderPath); os.IsNotExist(err) {
		os.MkdirAll(globalvariables.WorkflowLogFolderPath, 0777)
	}
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = globalvariables.WorkflowLogTimestampFormat
	formatter.FullTimestamp = globalvariables.WorkflowLogFullTimestamp
	log.SetFormatter(formatter)
	logFile, err := os.OpenFile(
		filepath.Join(globalvariables.WorkflowLogFolderPath, globalvariables.WorkflowLogFileName),
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		globalvariables.WorkflowLogFilePermissionCode,
	)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
}

var kafkaReader = kafkamanager.GetKafkaReader()
var ctx = context.Background()

func main() {
	setUpLogger()
	for {
		data, err := kafkaReader.ReadMessages(ctx)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("Data : ", string(data))
	}
}
