package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
)

func EditUserRequest(r *http.Request) (dto.EditUserDTO, error) {
	var (
		editUserDTO dto.EditUserDTO
		err         = json.NewDecoder(r.Body).Decode(&editUserDTO)
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	err = validate.Struct(editUserDTO)

	return editUserDTO, err
}
