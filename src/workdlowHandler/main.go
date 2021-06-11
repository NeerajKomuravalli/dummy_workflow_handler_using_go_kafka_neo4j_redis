package main

import (
	"os"
	"path/filepath"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
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

var kafkaListner = NewKafkaListner()

func main() {
	setUpLogger()
	go kafkaListner.Listen()
	c := 0
	for {
		c += 1
	}
}