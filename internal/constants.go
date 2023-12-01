package internal

import (
	"time"
)

const (
	ConnReadIdleTimeoutS  = 10 * time.Second
	ConnWriteIdleTimeoutS = 10 * time.Second
	LocalFilePath         = ".yt-transcription-translation"
)
