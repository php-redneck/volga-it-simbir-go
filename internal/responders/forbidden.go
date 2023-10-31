package responders

import "net/http"

func Forbidden(w http.ResponseWriter) {
	if err := render.String(w, 403, ""); err != nil {
		panic(err)
	}
}
