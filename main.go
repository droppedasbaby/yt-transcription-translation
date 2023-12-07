package main

import (
	"context"
	"sync"

	"github.com/GrewalAS/yt-transcription-translation/internal"
	"github.com/GrewalAS/yt-transcription-translation/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		internal.PanicIfError(err, logger)
	}
	defer logger.Sync()

	wg := &sync.WaitGroup{}
	orchestrateServer(wg, logger)
	wg.Wait()
}

func orchestrateServer(wg *sync.WaitGroup, parentLogger *zap.Logger) {
	logger := parentLogger.With(zap.String("component", "cmd.orchestrateServer"))
	ctx, cancel := context.WithCancel(context.Background())

	logger.Info("Initializing db connection...")
	logger.Info("Starting server...")
	s := server.NewServer(ctx, logger)
	wg.Add(1)
	go func() {
		defer wg.Done()
		shErr := internal.WaitForShutdown(ctx, logger)
		internal.PanicIfError(shErr, logger)
		cancel()
	}()

	wg.Add(1)
	defer wg.Done()
	s.ManagerServerLifecycle(ctx)
}
