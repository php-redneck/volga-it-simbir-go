package middleware

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func Global() []func(handler http.Handler) http.Handler {
	return []func(h http.Handler) http.Handler{
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
		middleware.StripSlashes,
		middleware.Timeout(60 * time.Second),
	}
}
