package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
)

func CreateTransportRequest(r *http.Request) (dto.TransportDTO, error) {
	var (
		transportDto dto.TransportDTO
		err          = json.NewDecoder(r.Body).Decode(&transportDto)
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	err = validate.Struct(transportDto)

	return transportDto, err
}
