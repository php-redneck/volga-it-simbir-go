package dto

type PaginationDTO struct {
	Start int `json:"start" validate:"number,min=0"`
	Count int `json:"count" validate:"required,number,min=1"`
}
