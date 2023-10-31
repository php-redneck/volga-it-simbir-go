package responders

import "net/http"

func Unauthorized(w http.ResponseWriter) {
	w.WriteHeader(401)
}
