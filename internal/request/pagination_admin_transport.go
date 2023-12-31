package request

import (
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
	"strconv"
)

func PaginationAdminTransportRequest(r *http.Request) (dto.TransportPaginationDTO, error) {
	start, _ := strconv.Atoi(r.URL.Query().Get("start"))

	count, _ := strconv.Atoi(r.URL.Query().Get("count"))

	transportType := r.URL.Query().Get("transportType")

	paginationDTO := dto.TransportPaginationDTO{
		Start:         start,
		Count:         count,
		TransportType: transportType,
	}

	validate := validator.New()

	err := validate.Struct(paginationDTO)

	return paginationDTO, err
}
