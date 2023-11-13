package internal

import (
	"time"
)

const (
	ConnReadIdleTimeoutS  = 10 * time.Second
	ConnWriteIdleTimeoutS = 10 * time.Second
	LocalFilePath         = ".yt-transcription-translation"
	DBFileName            = "db.sqlite"
	RunIDKey              = "run_id"
	EntClientKey          = "ent_client"
)
