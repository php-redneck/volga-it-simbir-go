package responders

import (
	"github.com/charmbracelet/log"
	"net/http"
)

func NoContent(w http.ResponseWriter) {
	if err := render.NoContent(w); err != nil {
		log.Error(err)
	}
}
