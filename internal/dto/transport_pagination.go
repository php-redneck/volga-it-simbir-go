package dto

type TransportPaginationDTO struct {
	Start         int    `json:"start" validate:"number,min=0"`
	Count         int    `json:"count" validate:"required,number,min=1"`
	TransportType string `json:"transportType" validate:"required,oneof=Car Bike Scooter All"`
}
