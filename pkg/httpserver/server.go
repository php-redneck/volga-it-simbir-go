package httpserver

import (
	"context"
	"errors"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/php-redneck/volga-it-simbir-go/internal/config"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	server http.Server
	Closer *Closer
}

func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		server: http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: handler,
		},
		Closer: &Closer{},
	}
}

func (s *Server) Run() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	s.Closer.Add(s.server.Shutdown)

	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server startup error: %v", err)
		}
	}()

	log.Infof("server started on http://127.0.0.1%s", s.server.Addr)

	log.Infof("swagger started on http://127.0.0.1%s%s/index.html", s.server.Addr, config.Swagger().BasePathPrefix)

	<-ctx.Done()

	log.Infof("shutting down server gracefully")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)

	defer cancel()

	if err := s.Closer.Close(shutdownCtx); err != nil {
		log.Error(err.Error())
	}
}
