package http

import (
	"errors"
	"github.com/eclipsemode/go-todo-app/internal/config"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg *config.Config, r *gin.Engine) *Server {
	return &Server{
		srv: &http.Server{
			Addr:         cfg.HTTPServer.Addr,
			Handler:      r,
			ReadTimeout:  cfg.HTTPServer.Timeout,
			WriteTimeout: cfg.HTTPServer.Timeout,
			IdleTimeout:  cfg.HTTPServer.Timeout,
		},
	}
}

func (s *Server) Start(log *slog.Logger) {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("error starting server", slog.Any("config", s.srv))
		}
	}()
}
