package dto

type EditUserDTO struct {
	Username string  `json:"username" validate:"omitempty,min=1,max=255"`
	Password string  `json:"password" validate:"omitempty,min=1,max=255"`
	IsAdmin  bool    `json:"isAdmin" validate:"boolean"`
	Balance  float32 `json:"balance" validate:"number"`
}
