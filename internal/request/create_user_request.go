package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
)

func CreateUserRequest(r *http.Request) (dto.CreateUserDTO, error) {
	var (
		createUserDTO dto.CreateUserDTO
		err           = json.NewDecoder(r.Body).Decode(&createUserDTO)
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	err = validate.Struct(createUserDTO)

	return createUserDTO, err
}
