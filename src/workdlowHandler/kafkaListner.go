package main

import (
	"context"
	"fmt"

	kafkamanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/kafkaManager"
	log "github.com/sirupsen/logrus"
)

type KafkaListner struct {
	Context context.Context
}

func NewKafkaListner() *KafkaListner {
	return &KafkaListner{
		context.Background(),
	}
}

func (kl *KafkaListner) Listen() {
	kafkaReader := kafkamanager.GetKafkaReader()
	for {
		data, err := kafkaReader.ReadMessages(kl.Context)
		if err != nil {
			log.Error(err)
		}
		fmt.Println("Data : ", string(data))
	}
}
