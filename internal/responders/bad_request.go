package responders

import (
	"github.com/charmbracelet/log"
	"net/http"
)

type BadRequestResponse struct {
	Errors []string `json:"errors"`
}

func BadRequest(w http.ResponseWriter, errors []string) {
	err := render.JSON(w, 400, BadRequestResponse{errors})

	if err != nil {
		log.Error(err.Error())
	}
}
