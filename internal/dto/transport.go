package dto

import "github.com/guregu/null"

type TransportDTO struct {
	CanBeRented   bool        `json:"canBeRented" validate:"required,boolean"`
	TransportType string      `json:"transportType" validate:"required,oneof=Car Bike Scooter"`
	Model         string      `json:"model" validate:"required"`
	Color         string      `json:"color" validate:"required"`
	Identifier    string      `json:"identifier" validate:"required"`
	Description   null.String `json:"description" validate:"omitempty" swaggertype:"string"`
	Latitude      float32     `json:"latitude" validate:"required,latitude"`
	Longitude     float32     `json:"longitude" validate:"required,longitude"`
	MinutePrice   null.Float  `json:"minutePrice" validate:"omitempty,number" swaggertype:"number"`
	DayPrice      null.Float  `json:"dayPrice" validate:"omitempty,number" swaggertype:"number"`
}
