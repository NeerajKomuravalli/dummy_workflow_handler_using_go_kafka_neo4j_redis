package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/server/errors"
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/server/models"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

func addDevice(w http.ResponseWriter, r *http.Request) {
	deviceData := models.DeviceData{}
	err := json.NewDecoder(r.Body).Decode(&deviceData)
	if err != nil {
		errors.ManageErrorResponse(w, fmt.Sprint(err), errors.JsonDecoderError)
		return
	}
	err = validate.Struct(&deviceData)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
		errors.ManageErrorResponse(w, fmt.Sprint(err), errors.ValidationError)
		return
	}
	// Put rest of the logic here
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
	w.Write([]byte(fmt.Sprintf("/getDevice %s\n", params["id"])))
}

var validate *validator.Validate

func main() {
	validate = validator.New()

	router := mux.NewRouter()
	router.HandleFunc("/addDevice", addDevice).Methods("POST")
	router.HandleFunc("/updateDeviceProperty", updateDeviceProperty).Methods("PUT")
	router.HandleFunc("/getDevice/{id}", getDevice).Methods("GET")
	http.ListenAndServe(":8080", router)
}
