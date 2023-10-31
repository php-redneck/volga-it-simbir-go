package entities

type User struct {
	Model
	Username string  `gorm:"type:varchar;size:255" json:"username"`
	Password string  `gorm:"type:varchar;size:255" json:"-"`
	IsAdmin  bool    `json:"isAdmin"`
	Balance  float32 `json:"balance"`
}
