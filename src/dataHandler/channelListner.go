package main

import (
	"context"
	"fmt"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
	log "github.com/sirupsen/logrus"
)

type ChannelListner struct {
	Channel chan models.DeviceDataPair
	Context context.Context
}

func NewChannelListner() *ChannelListner {
	return &ChannelListner{
		make(chan models.DeviceDataPair, globalvariables.DataHandlerBufferSize),
		ctx,
	}
}

func (cl *ChannelListner) ListenAndAddToKafka() {
	for {
		dataPair := <-cl.Channel
		fmt.Println("Output : ", dataPair.DeviceData)
		go manageKafkaWriteProcess(cl.Context, dataPair)
	}
}

func manageKafkaWriteProcess(channelListnerContex context.Context, dataPair models.DeviceDataPair) {
	err := kafkaWriter.WriteMessages(
		channelListnerContex,
		[]byte(dataPair.DeviceData.Id),
		[]byte(dataPair.JsonData),
	)
	if err != nil {
		log.Error(err)
		fmt.Println("error : ", err)
		err = serverRedisClient.PutData(channelListnerContex, dataPair.DeviceData.Id, dataPair.DeviceData, 0)
		if err != nil {
			fmt.Println("Redis put data error : ", err)
			log.Error(err)
		}
	}
}
