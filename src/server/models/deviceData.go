package models

type DeviceData struct {
	Name        string           `json:"name" validate:"required"`
	Id          string           `json:"id" validate:"required"`
	Description string           `json:"description"`
	Properties  DeviceProperties `json:"properties" validate:"required"`
}

type DeviceProperties struct {
	Status     string  `json:"status" validate:"required"`
	Temprature float64 `json:"temprature" validate:"required"`
}
