package main

import (
	"context"
	"encoding/json"
	"fmt"

	kafkamanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/kafkaManager"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
	log "github.com/sirupsen/logrus"
)

type KafkaListner struct {
	Context    context.Context
	OutChannel ChannelListner
}

func NewKafkaListner(ctx context.Context, outChannel ChannelListner) *KafkaListner {
	return &KafkaListner{
		context.Background(),
		outChannel,
	}
}

func (kl *KafkaListner) Listen() {
	kafkaReader := kafkamanager.GetKafkaReader()
	for {
		data, err := kafkaReader.ReadMessages(kl.Context)
		if err != nil {
			log.Error(err)
		}
		deviceData := models.DeviceData{}
		json.Unmarshal(data, &deviceData)
		fmt.Println("Data : ", deviceData)
		log.Info(fmt.Sprintln("Received data : ", deviceData))
		kl.OutChannel.Channel <- deviceData
	}
}
