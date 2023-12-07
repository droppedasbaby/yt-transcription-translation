package server

import (
	"context"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/GrewalAS/yt-transcription-translation/internal"
)

type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
	dm         *internal.DownloadManager
}

func NewServer(ctx context.Context, parentLogger *zap.Logger) *Server {
	openaiKey, err := internal.GetEnvVars("OPENAI_KEY")
	if err != nil {
		panic("Could not get OpenAI key")
	}

	logger := parentLogger.With(zap.String("component", "Server"))
	handler := http.NewServeMux()
	httpServer := &http.Server{
		Addr:         ":61235",
		Handler:      handler,
		ReadTimeout:  internal.ConnReadIdleTimeoutS,
		WriteTimeout: internal.ConnWriteIdleTimeoutS,
	}
	server := &Server{
		logger:     logger,
		httpServer: httpServer,
		dm:         internal.NewDownloadManager(logger, openaiKey, ctx),
	}
	server.configureRoutes(handler)
	return server
}

func (s *Server) configureRoutes(handler *http.ServeMux) {
	handler.Handle("/start",
		internal.ChainMiddleware(
			http.HandlerFunc(s.startHandler),
			internal.MethodChecker(http.MethodPost),
			internal.JSONHeadersMiddleware(),
			internal.LoggingMiddleware(s.logger),
		),
	)
	handler.Handle("/results",
		internal.ChainMiddleware(
			http.HandlerFunc(s.resultsHandler),
			internal.MethodChecker(http.MethodPost),
			internal.JSONHeadersMiddleware(),
			internal.LoggingMiddleware(s.logger),
		),
	)
}

func (s *Server) ManagerServerLifecycle(ctx context.Context) {
	s.logger = s.logger.With(zap.String("component", "Server.ManagerServerLifecycle"))
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
