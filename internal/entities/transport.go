package entities

import (
	"github.com/guregu/null"
	"gorm.io/gorm"
	"time"
)

type Transport struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	OwnerId       int            `json:"ownerId"`
	CanBeRented   bool           `json:"canBeRented"`
	TransportType string         `gorm:"type:varchar;size:255" json:"transportType"`
	Model         string         `gorm:"type:varchar;size:255" json:"model"`
	Color         string         `gorm:"type:varchar;size:255" json:"color"`
	Identifier    string         `gorm:"type:varchar;size:255" json:"identifier"`
	Description   null.String    `gorm:"null" json:"description"`
	Latitude      float32        `json:"latitude"`
	Longitude     float32        `json:"longitude"`
	MinutePrice   null.Float     `gorm:"null" json:"minutePrice"`
	DayPrice      null.Float     `gorm:"null" json:"dayPrice"`
}
