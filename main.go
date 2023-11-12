package main

import (
	"context"
	"sync"

	"entgo.io/ent/dialect"
	"go.uber.org/zap"

	"github.com/GrewalAS/yt-transcription-translation/ent"
	"github.com/GrewalAS/yt-transcription-translation/internal"
	"github.com/GrewalAS/yt-transcription-translation/server"
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
	dbPath, err := internal.CreateDirAndGetFullPath(internal.DBFileName)
	if err != nil {
		internal.PanicIfError(err, logger)
	}
	client, err := ent.Open(dialect.SQLite, dbPath+"?_fk=1")
	if err != nil {
		internal.PanicIfError(err, logger)
	}
	if err = client.Schema.Create(ctx); err != nil {
		internal.PanicIfError(err, logger)
	}

	logger.Info("Starting server...")
	s := server.NewServer(client, logger)
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
