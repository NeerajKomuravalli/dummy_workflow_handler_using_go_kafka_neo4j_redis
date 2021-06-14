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
	// Logging
	KafkaLogFolderPath         = "logs"
	KafkaLogFileName           = "kafka_logs.txt"
	KafkaLogTimestampFormat    = "02-01-2006 15:04:05"
	KafkaLogFullTimestamp      = true
	KafkaLogFilePermissionCode = 0644

	// Workflow handler
	WorkflowBufferSize = 10
	// Logging
	WorkflowLogFolderPath         = "logs"
	WorkflowLogFileName           = "logs.txt"
	WorkflowLogTimestampFormat    = "02-01-2006 15:04:05"
	WorkflowLogFullTimestamp      = true
	WorkflowLogFilePermissionCode = 0644
)
