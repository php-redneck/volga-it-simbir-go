package dto

type AdminTransportDTO struct {
	OwnerId       int     `json:"ownerId" validation:"required,number"`
	CanBeRented   bool    `json:"canBeRented" validation:"required,boolean"`
	TransportType string  `json:"transportType" validation:"required,oneof=Car Bike Scooter"`
	Model         string  `json:"model" validation:"required"`
	Color         string  `json:"color" validation:"required"`
	Identifier    string  `json:"identifier" validation:"required"`
	Description   string  `json:"description"`
	Latitude      float32 `json:"latitude" validation:"required,latitude"`
	Longitude     float32 `json:"longitude" validation:"required,longitude"`
	MinutePrice   float32 `json:"minutePrice" validation:"number"`
	DayPrice      float32 `json:"dayPrice" validation:"number"`
}
