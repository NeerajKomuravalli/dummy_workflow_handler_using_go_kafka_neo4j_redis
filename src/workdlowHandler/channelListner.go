package main

import (
	"context"
	"fmt"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
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
}
