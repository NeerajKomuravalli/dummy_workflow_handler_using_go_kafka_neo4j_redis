package main

import (
	"context"
	"fmt"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type ChannelListner struct {
	Channel chan models.DeviceData
	Context context.Context
}

func NewChannelListner(ctx context.Context) *ChannelListner {
	return &ChannelListner{
		make(chan models.DeviceData, globalvariables.WorkflowBufferSize),
		ctx,
	}
}

func (cl *ChannelListner) ListenAndTriggerWorkFlow() {
	for {
		deviceData := <-cl.Channel
		log.Info(fmt.Sprintln("Channel received the data ", deviceData))
		go triggerWorkFlow(deviceData)
	}
}

func triggerWorkFlow(deviceData models.DeviceData) {
	// All workflow related code will reside here for now neo4j code will be triggered from here
	fmt.Println("Data at workflow : ", deviceData)
	fn := GenerateDeviceDataNodeCreationFn(deviceData)
	session := newNeo4jClient.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close()
	result, err := session.WriteTransaction(fn)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Result : ", result)
	log.Info(fmt.Sprintln("Added this device to neo4j : ", result))
}
