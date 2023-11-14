package internal

import (
	"context"

	"go.uber.org/zap"

	"github.com/GrewalAS/yt-transcription-translation/ent"
	"github.com/GrewalAS/yt-transcription-translation/ent/record"
)

func ProcessVideoID(ctx context.Context, parentLogger *zap.Logger, rec *ent.Record) {
	logger := parentLogger.With(zap.String("component", "internal.ProcessVideoID"))

	if rec.Status == record.StatusDownloading {
		fileLoc, _, err := DownloadVideo(rec.VideoID)
		if err != nil {
			rec, err = rec.Update().SetStatus(record.StatusError).Save(ctx)
			if err != nil {
				logger.Error("failed to update record when download failed", zap.Error(err))
			}
			logger.Error("failed to download video", zap.Error(err))
		}

		// TODO: change to updating the status
		_, err = rec.Update().
			SetFileLocation(fileLoc).
			SetStatus(record.StatusDownloaded).
			Save(ctx)
		if err != nil {
			logger.Error("failed to update record when download succeeded", zap.Error(err))
		}
	}

	// if rec.Status == record.StatusDownloaded {
	// }
	// if rec.Status == record.StatusTranscribed {
	// }
}
