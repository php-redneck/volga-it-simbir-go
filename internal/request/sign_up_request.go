package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
)

func SignUpRequest(r *http.Request) (dto.SignUpDTO, error) {
	var (
		signUpDTO dto.SignUpDTO
		err       = json.NewDecoder(r.Body).Decode(&signUpDTO)
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	err = validate.Struct(signUpDTO)

	return signUpDTO, err
}
