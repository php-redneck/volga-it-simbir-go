package request

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/php-redneck/volga-it-simbir-go/internal/dto"
	"net/http"
)

func SignInRequest(r *http.Request) (dto.SignInDTO, error) {
	var (
		signInDTO dto.SignInDTO
		err       = json.NewDecoder(r.Body).Decode(&signInDTO)
	)

	if err != nil {
		panic(err)
	}

	validate := validator.New()

	err = validate.Struct(signInDTO)

	return signInDTO, err
}
