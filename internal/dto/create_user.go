package dto

type CreateUserDTO struct {
	Username string  `json:"username" validate:"required,min=1,max=255"`
	Password string  `json:"password" validate:"required,min=1,max=255"`
	IsAdmin  bool    `json:"isAdmin" validate:"boolean"`
	Balance  float32 `json:"balance" validate:"number"`
}
