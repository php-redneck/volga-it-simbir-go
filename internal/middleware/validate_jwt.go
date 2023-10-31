package middleware

import (
	"context"
	"github.com/php-redneck/volga-it-simbir-go/internal/responders"
	"github.com/php-redneck/volga-it-simbir-go/internal/services"
	"net/http"
)

func ValidateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		if authToken == "" {
			responders.Unauthorized(rw)
			return
		}

		jwtService := services.JWTService{}

		if !jwtService.IsValid(authToken) {
			responders.Unauthorized(rw)
			return
		}

		tokenPayload := jwtService.Decode(authToken)

		isLogout := services.LogoutHistoryService{}.IsLogout(tokenPayload.JWTID)

		if isLogout {
			responders.Unauthorized(rw)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", tokenPayload.Id)
		ctx = context.WithValue(ctx, "tokenId", tokenPayload.JWTID)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
