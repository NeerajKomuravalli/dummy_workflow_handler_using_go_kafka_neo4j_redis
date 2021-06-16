package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	globalvariables "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/globalVariables"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
	neo4jmanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/neo4jManager"
	redismanager "github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/redisManager"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/server/errors"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func setUpLogger() {
	if _, err := os.Stat(globalvariables.ServerLogFolderPath); os.IsNotExist(err) {
		os.MkdirAll(globalvariables.ServerLogFolderPath, 0777)
	}
	formatter := new(log.TextFormatter)
	formatter.TimestampFormat = globalvariables.ServerLogTimestampFormat
	formatter.FullTimestamp = globalvariables.ServerLogFullTimestamp
	log.SetFormatter(formatter)
	logFile, err := os.OpenFile(
		filepath.Join(globalvariables.ServerLogFolderPath, globalvariables.ServerLogFileName),
		os.O_WRONLY|os.O_APPEND|os.O_CREATE,
		globalvariables.ServerLogFilePermissionCode,
	)
	if err != nil {
		log.Panic(err)
	}
	log.SetLevel(log.InfoLevel)
	log.SetOutput(logFile)
}

func addDevice(w http.ResponseWriter, r *http.Request) {
	deviceData := models.DeviceData{}
	err := json.NewDecoder(r.Body).Decode(&deviceData)
	if err != nil {
		log.Error(err)
		errors.ManageErrorResponse(w, fmt.Sprint(err), errors.JsonDecoderError)
		return
	}
	err = validate.Struct(&deviceData)
	if err != nil {
		log.Error(err)
		errors.ManageErrorResponse(w, fmt.Sprint(err), errors.ValidationError)
		return
	}
	// Put data on redis
	err = redisClient.PutData(ctx, deviceData.Id, deviceData, 0)
	if err != nil {
		log.Error(err)
		errors.ManageErrorResponse(w, fmt.Sprint(err), errors.RedisPutDataError)
		return
	}

	log.Info(fmt.Sprintln("Successfully sending response with data : ", deviceData))
	successResp := models.SuccessRequest{
		Success:    true,
		DeviceData: deviceData,
	}
	json.NewEncoder(w).Encode(successResp)
}

func updateDeviceProperty(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/updateDeviceProperty"))
}

func getDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //params["id"] => has id info
	fn := neo4jmanager.GenerateGetDeviceDataNodeFn(params["id"])
	session := newNeo4jClient.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()
	deviceData, err := session.ReadTransaction(fn)
	if err != nil {
		log.Error(err)
		errors.ManageErrorResponse(w, fmt.Sprint(err), errors.Neo4jReadTransactionError)
		return
	}
	log.Info(fmt.Sprintln("Successfully sending response with data (getDevice) : ", deviceData.(models.DeviceData)))
	successResp := models.SuccessRequest{
		Success:    true,
		DeviceData: deviceData.(models.DeviceData),
	}
	json.NewEncoder(w).Encode(successResp)
}

var validate *validator.Validate
var ctx = context.Background()
var redisClient = redismanager.GetRedisClient(
	globalvariables.ServerRedisIpAddress,
	globalvariables.ServerRedisPassword,
	globalvariables.ServerRedisDbIndex,
)
var newNeo4jClient = neo4jmanager.NewNeo4jClient(
	globalvariables.Neo4jUri,
	globalvariables.Neo4jDbName,
	globalvariables.Neo4jDbPassword,
)

func main() {
	setUpLogger()
	validate = validator.New()

	router := mux.NewRouter()
	router.HandleFunc("/addDevice", addDevice).Methods("POST")
	router.HandleFunc("/updateDeviceProperty", updateDeviceProperty).Methods("PUT")
	router.HandleFunc("/getDevice/{id}", getDevice).Methods("GET")
	http.ListenAndServe(":8080", router)
}
