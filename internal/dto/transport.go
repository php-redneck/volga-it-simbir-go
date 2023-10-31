package dto

import "github.com/guregu/null"

type TransportDTO struct {
	CanBeRented   bool        `json:"canBeRented" validation:"required,boolean"`
	TransportType string      `json:"transportType" validation:"required,oneof=Car Bike Scooter"`
	Model         string      `json:"model" validation:"required"`
	Color         string      `json:"color" validation:"required"`
	Identifier    string      `json:"identifier" validation:"required"`
	Description   null.String `json:"description"`
	Latitude      float32     `json:"latitude" validation:"required,latitude"`
	Longitude     float32     `json:"longitude" validation:"required,longitude"`
	MinutePrice   null.Float  `json:"minutePrice" validation:"number"`
	DayPrice      null.Float  `json:"dayPrice" validation:"number"`
}
