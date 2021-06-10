package models

type SuccessRequest struct {
	Success    bool       `json:"success" validate:"required"`
	DeviceData DeviceData `json:"data" validate:"required"`
}

type FailureRequest struct {
	Success bool  `json:"success" validate:"required"`
	Error   Error `json:"error" validate:"required"`
}

type Error struct {
	ErrorCode    int    `json:"errorCode" validate:"required"`
	ErrorMessage string `json:"errorMessage" validate:"required"`
}
