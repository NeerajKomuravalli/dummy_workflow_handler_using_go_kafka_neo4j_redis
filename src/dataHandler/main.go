package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
	redismanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/redisManager"
)

// Data handler

// - Poll on redis for updates
// - filter and pre-process data and put it on a Kafka bus

func pollOnRedis() {
	for {
		iter := serverRedisClient.Client.Scan(ctx, 0, "*", 0).Iterator()
		for iter.Next(ctx) {
			// fmt.Println(iter.Val()) // Will print keys one after another
			handleData(iter.Val())
		}
		if err := iter.Err(); err != nil {
			panic(err)
		}
		time.Sleep(PollDeration * time.Millisecond)
	}
}

func handleData(key string) {
	data, err := serverRedisClient.GetData(ctx, key)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	deviceData := models.DeviceData{}
	json.Unmarshal([]byte(data), &deviceData)
	deviceDataChannel <- deviceData
	serverRedisClient.DeleteData(ctx, key)
}

func channelListner() {
	for {
		deviceData := <-deviceDataChannel
		fmt.Println("Output : ", deviceData)
		err := addToKafka()
		if err != nil {
			serverRedisClient.PutData(ctx, deviceData.Id, deviceData, 0)
		}
	}
}

func addToKafka() error {
	fmt.Println("Add to kafka")
	// Add data to kafka and send error if unsuccessfull
	err := fmt.Errorf("Error")
	return err
}

// To get the data from the server
var serverRedisClient = redismanager.GetRedisClient(
	globalvariables.ServerRedisIpAddress,
	globalvariables.ServerRedisPassword,
	globalvariables.ServerRedisDbIndex,
)
var ctx = context.Background()
var deviceDataChannel = make(chan models.DeviceData, globalvariables.DataHandlerBufferSize)

func main() {
	go channelListner()
	pollOnRedis()
}
