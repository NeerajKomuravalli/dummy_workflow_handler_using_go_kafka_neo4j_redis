package errors

import (
	"encoding/json"
	"net/http"

	"github.com/NeerajKomuravalli/dummy_workflow_handler_using_go_kafka_neo4j_redis/src/server/models"
)

func ManageErrorResponse(w http.ResponseWriter, err string, errCode int) {
	errData := models.FailureRequest{
		Success: false,
		Error: models.Error{
			ErrorCode:    errCode,
			ErrorMessage: err,
		},
	}
	json.NewEncoder(w).Encode(&errData)
}
