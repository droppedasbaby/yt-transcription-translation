package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	ConnReadIdleTimeoutS  = 10 * time.Second
	ConnWriteIdleTimeoutS = 10 * time.Second
	LocalFilePath         = ".yt-transcription-translation"
	DBFileName            = "db.sqlite"
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
