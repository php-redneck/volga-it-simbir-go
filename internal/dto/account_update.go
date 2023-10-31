package dto

type AccountUpdateDTO struct {
	Username string `json:"username" validate:"omitempty,min=1,max=255"`
	Password string `json:"password" validate:"omitempty,min=1,max=255"`
}
