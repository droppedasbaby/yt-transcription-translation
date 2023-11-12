package internal

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

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

func MethodChecker(h http.HandlerFunc, allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, method := range allowedMethods {
			if r.Method == method {
				h(w, r)
				return
			}
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
