package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
)

func EditTransportRequest(r *http.Request) (dto.TransportDTO, error) {
	var (
		editTransportDTO dto.TransportDTO
		err              = json.NewDecoder(r.Body).Decode(&editTransportDTO)
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	err = validate.Struct(editTransportDTO)

	return editTransportDTO, err
}
