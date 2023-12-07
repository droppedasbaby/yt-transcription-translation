package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func CreateDirAndGetFullPath(filename string) (string, error) {
	expandedPath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("get home directory: %w", err)
	}

	fullPath := filepath.Join(expandedPath, LocalFilePath)
	err = os.MkdirAll(fullPath, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("create directory: %w", err)
	}

	return filepath.Join(fullPath, filename), nil
}

func WaitForShutdown(ctx context.Context, logger *zap.Logger) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)

	select {
	case <-c:
		logger.Info("Received SIGINT, no longer waiting for shutdown.")
		return nil
	case <-ctx.Done():
		logger.Info("Context cancelled, no longer waiting for shutdown.")
		return fmt.Errorf("WaitForShutdown: context cancelled")
	}
}

func PanicIfError(err error, logger *zap.Logger) {
	if err != nil {
		logger.Fatal("PanicIfError", zap.Error(err))
		panic(err)
	}
}

type ChainMiddlewareFunc func(http.Handler) http.Handler

func ChainMiddleware(h http.Handler, middlewares ...ChainMiddlewareFunc) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func MethodChecker(allowedMethods ...string) ChainMiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, method := range allowedMethods {
				if r.Method == method {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		})
	}
}

func LoggingMiddleware(logger *zap.Logger) ChainMiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.New()
			logger.Info(
				"Request received",
				zap.String("request_id", requestID.String()),
				zap.String("remote_addr", r.RemoteAddr),
				zap.String("host", r.Host),
				zap.String("proto", r.Proto),
				zap.String("method", r.Method),
				zap.String("url", r.URL.String()),
			)
			next.ServeHTTP(w, r)
			logger.Info("Request completed", zap.String("request_id", requestID.String()))
		})
	}
}

func JSONHeadersMiddleware() ChainMiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
}

func GetEnvVars(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("GetEnvVars: %s not set", key)
	}
	return value, nil
}
