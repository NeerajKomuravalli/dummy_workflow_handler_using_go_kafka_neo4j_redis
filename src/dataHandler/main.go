package main

import (
	"context"
	"fmt"
	"time"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	redismanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/redisManager"
)

// Data handler

// - Poll on redis for updates
// - filter and pre-process data and put it on a Kafka bus

func pollOnRedis() {
	for {
		iter := redisClient.Client.Scan(ctx, 0, "*", 0).Iterator()
		for iter.Next(ctx) {
			fmt.Println(iter.Val()) // Will print keys one after another
		}
		if err := iter.Err(); err != nil {
			panic(err)
		}
		time.Sleep(PollDeration * time.Millisecond)
	}
}

var redisClient = redismanager.GetRedisClient(
	globalvariables.ServerRedisIpAddress,
	globalvariables.ServerRedisPassword,
	globalvariables.ServerRedisDbIndex,
)
var ctx = context.Background()

func main() {
	pollOnRedis()
}
