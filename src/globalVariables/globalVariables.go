package globalvariables

const (
	// Server
	ServerRedisIpAddress = "localhost:6379"
	ServerRedisPassword  = ""
	ServerRedisDbIndex   = 0

	// Data handler
	DataHandlerPollDeration   = 10 // In milliseconds
	DataHandlerRedisIpAddress = "localhost:6380"
	DataHandlerRedisPassword  = ""
	DataHandlerRedisDbIndex   = 0
	DataHandlerBufferSize     = 10
	// Logging
	DataHandlerLogFolderPath         = "logs"
	DataHandlerLogFileName           = "logs.txt"
	DataHandlerLogTimestampFormat    = "02-01-2006 15:04:05"
	DataHandlerLogFullTimestamp      = true
	DataHandlerLogFilePermissionCode = 0644

	// Kafka
	KafkaBorkerAddress = "localhost:9092"
	KafkaTopicName     = "messenger"
	KafkaGroupId       = "my-group"
)
