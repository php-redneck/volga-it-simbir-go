package responders

import (
	"github.com/charmbracelet/log"
	"net/http"
)

func OK(w http.ResponseWriter, payload interface{}) {
	err := render.JSON(w, 200, payload)

	if err != nil {
		log.Error(err.Error())
	}
}
