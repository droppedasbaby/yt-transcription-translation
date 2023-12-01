package internal_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"syscall"
	"testing"
	"time"

	"github.com/GrewalAS/yt-transcription-translation/internal"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateDirAndGetFullPath(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "test")
	originalHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer func() {
		os.Setenv("HOME", originalHome)
		os.RemoveAll(tmpDir)
	}()

	filename := "test.txt"
	path, err := internal.CreateDirAndGetFullPath(filename)
	assert.NoError(t, err)
	assert.Contains(t, path, filename)
}

func TestWaitForShutdown(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		cancel()
	}()
	err := internal.WaitForShutdown(ctx, logger)
	assert.Error(t, err)

	// Test with SIGINT signal
	ctx, cancel = context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * 100)
		syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()
	err = internal.WaitForShutdown(ctx, logger)
	cancel()
	assert.NoError(t, err)
}

func TestChainMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	middleware := internal.LoggingMiddleware(zap.NewNop())
	chained := internal.ChainMiddleware(handler, middleware)
	assert.NotNil(t, chained)
}

func TestMethodChecker(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middleware := internal.MethodChecker("GET")
	chained := middleware(handler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	chained.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test with unallowed method
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/", nil)
	chained.ServeHTTP(w, req)
	assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
}

func TestJSONHeadersMiddleware(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	middleware := internal.JSONHeadersMiddleware()
	chained := middleware(handler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	chained.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
