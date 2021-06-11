package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	kafkamanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/kafkaManager"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
	redismanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/redisManager"
)

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
	// We unmarshel only to make sure the data being sent is right
	err = json.Unmarshal([]byte(data), &deviceData)
	if err != nil {
		// This is the case when the data coming into redis is not valid and we need to filter it out
		fmt.Println("Handle this case! :: Filter this data out")
	}
	dataPair := deviceDataPair{
		data,
		deviceData,
	}
	deviceDataChannel <- dataPair
	serverRedisClient.DeleteData(ctx, key)
}

func channelListner() {
	for {
		dataPair := <-deviceDataChannel
		fmt.Println("Output : ", dataPair.DeviceData)
		err := addToKafka(dataPair)
		if err != nil {
			fmt.Println("error : ", err)
			err = serverRedisClient.PutData(ctx, dataPair.DeviceData.Id, dataPair.DeviceData, 0)
			if err != nil {
				fmt.Println("Redis put data error : ", err)
			}
			// break
		}
	}
}

func addToKafka(dataPair deviceDataPair) error {
	fmt.Printf("Add to kafka : %s\n", dataPair.JsonData)
	// Add data to kafka and send error if unsuccessfull
	err := kafkaWriter.WriteMessages(ctx, []byte(dataPair.DeviceData.Id), []byte(dataPair.JsonData))
	return err
}

type deviceDataPair struct {
	JsonData   string
	DeviceData models.DeviceData
}

// To get the data from the server
var serverRedisClient = redismanager.GetRedisClient(
	globalvariables.ServerRedisIpAddress,
	globalvariables.ServerRedisPassword,
	globalvariables.ServerRedisDbIndex,
)
var ctx = context.Background()
var deviceDataChannel = make(chan deviceDataPair, globalvariables.DataHandlerBufferSize)

// Get kafka writer
var kafkaWriter = kafkamanager.GetKafkaWriter()

func main() {
	go channelListner()
	pollOnRedis()
}
