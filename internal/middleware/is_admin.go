package middleware

import (
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
)

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		user, err := services.UserService{}.Show(r.Context().Value("userId").(int))

		if err != nil || !user.IsAdmin {
			responders.Forbidden(rw)
			return
		}

		next.ServeHTTP(rw, r)
	})
}
