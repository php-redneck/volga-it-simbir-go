package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
	"strconv"
)

func PaginationRequest(r *http.Request) (dto.PaginationDTO, error) {
	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	count, _ := strconv.Atoi(r.URL.Query().Get("count"))

	paginationDTO := dto.PaginationDTO{
		Start: start,
		Count: count,
	}

	validate := validator.New()

	err := validate.Struct(paginationDTO)

	return paginationDTO, err
}
