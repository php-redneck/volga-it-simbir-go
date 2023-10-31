package responders

import "net/http"

func NotFound(w http.ResponseWriter) {
	err := render.String(w, 404, "")

	if err != nil {
		panic(err)
	}
}
