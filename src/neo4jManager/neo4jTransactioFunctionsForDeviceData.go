package neo4jmanager

import (
	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/models"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GenerateDeviceDataNodeCreationFn(deviceData models.DeviceData) neo4j.TransactionWork {
	generatedFunc := func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run(
			"MERGE (n:Device { id: $id, name: $name, description: $description, status: $status, temprature: $temprature}) RETURN n.id, n.name, n.description, n.status, n.temprature",
			map[string]interface{}{
				"name":        deviceData.Name,
				"id":          deviceData.Id,
				"description": deviceData.Description,
				"status":      deviceData.Properties.Status,
				"temprature":  deviceData.Properties.Temprature,
			},
		)
		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.
		if err != nil {
			return nil, err
		}

		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
		return &models.DeviceData{
			Id:          record.Values[0].(string),
			Name:        record.Values[1].(string),
			Description: record.Values[2].(string),
			Properties: models.DeviceProperties{
				Status:     record.Values[3].(string),
				Temprature: record.Values[4].(float64),
			},
		}, nil
	}

	return generatedFunc
}

func GenerateGetDeviceDataNodeFn(deviceDataId string) neo4j.TransactionWork {
	generatedFunc := func(tx neo4j.Transaction) (interface{}, error) {
		query := "MATCH (n:Device {id: $id}) return n.id, n.name, n.description, n.status, n.temprature"
		records, err := tx.Run(
			query,
			map[string]interface{}{
				"id": deviceDataId,
			},
		)
		// In face of driver native errors, make sure to return them directly.
		// Depending on the error, the driver may try to execute the function again.
		if err != nil {
			return nil, err
		}

		record, err := records.Single()
		if err != nil {
			return nil, err
		}
		// You can also retrieve values by name, with e.g. `id, found := record.Get("n.id")`
		return models.DeviceData{
			Id:          record.Values[0].(string),
			Name:        record.Values[1].(string),
			Description: record.Values[2].(string),
			Properties: models.DeviceProperties{
				Status:     record.Values[3].(string),
				Temprature: record.Values[4].(float64),
			},
		}, nil
	}

	return generatedFunc
}
