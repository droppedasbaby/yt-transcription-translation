package server

import (
	"context"
	"errors"
	"net/http"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"go.uber.org/zap"

	"github.com/GrewalAS/yt-transcription-translation/ent"
	"github.com/GrewalAS/yt-transcription-translation/internal"
)

type Server struct {
	client     *ent.Client
	httpServer *http.Server
	logger     *zap.Logger
}

func NewServer(client *ent.Client, parentLogger *zap.Logger) *Server {
	logger := parentLogger.With(zap.String("component", "cmd.Server"))
	handler := http.NewServeMux()
	handler.HandleFunc("/start", startHandler)

	server := &http.Server{
		Addr:         ":61235",
		Handler:      handler,
		ReadTimeout:  internal.ConnReadIdleTimeoutS,
		WriteTimeout: internal.ConnWriteIdleTimeoutS,
	}

	return &Server{client: client, httpServer: server, logger: logger}
}

func (s *Server) ManagerServerLifecycle(ctx context.Context) {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			if errors.Is(err, http.ErrServerClosed) {
				s.logger.Info("Server no longer listening")
			} else {
				s.logger.Fatal("Failed to start server", zap.Error(err))
			}
		}
	}()

	s.logger.Info("Server running...")
	<-ctx.Done()
	err := s.httpServer.Shutdown(ctx)

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("Error while shutting down server", zap.Error(err))
	}

	s.logger.Info("Server shutdown.")
}
