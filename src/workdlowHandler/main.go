package main

import (
	"context"
	"os"
	"path/filepath"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	neo4jmanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/neo4jManager"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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
	log.SetLevel(log.InfoLevel)
	log.SetOutput(logFile)
}

var ctx = context.Background()
var channelListner = NewChannelListner(ctx)
var kafkaListner = NewKafkaListner(ctx, *channelListner)
var newNeo4jClient = neo4jmanager.NewNeo4jClient("bolt://0.0.0.0:7687", "neo4j", "neo4j-testing")

type Item struct {
	Id   int64
	Name string
}

type MultipleNodes struct {
	Nodes []*neo4j.Record
}

func main() {
	setUpLogger()
	go kafkaListner.Listen()
	go channelListner.ListenAndTriggerWorkFlow()
	c := 0
	for {
		c += 1
	}
}
