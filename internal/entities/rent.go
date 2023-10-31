package entities

import (
	"github.com/guregu/null"
	"time"
)

type Rent struct {
	Model
	TransportId int        `json:"transportId"`
	UserId      int        `json:"userId"`
	TimeStart   time.Time  `json:"timeStart"`
	TimeEnd     null.Time  `json:"timeEnd" swaggertype:"string"`
	PriceOfUnit float32    `json:"priceOfUnit"`
	PriceType   string     `json:"priceType"`
	FinalPrice  null.Float `json:"finalPrice" swaggertype:"number"`
}
