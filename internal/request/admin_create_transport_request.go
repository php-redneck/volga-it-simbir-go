package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
)

func AdminCreateTransportRequest(r *http.Request) (dto.AdminTransportDTO, error) {
	var (
		transportDto dto.AdminTransportDTO
		err          = json.NewDecoder(r.Body).Decode(&transportDto)
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	err = validate.Struct(transportDto)

	return transportDto, err
}
