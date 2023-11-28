package server

// import (
// 	"context"
// 	"fmt"
//
// 	"github.com/google/uuid"
//
// 	"github.com/GrewalAS/yt-transcription-translation/ent"
// 	"github.com/GrewalAS/yt-transcription-translation/ent/record"
// )
//
// func FetchOrCreateRecordByVideoID(
// 	ctx context.Context,
// 	client *ent.Client,
// 	videoURL string,
// 	videoID string,
// 	runID uuid.UUID,
// ) (*ent.Record, error) {
// 	rec, err := client.Record.
// 		Query().
// 		Where(record.VideoID(videoID)).
// 		Only(ctx)
// 	if err != nil {
// 		if !ent.IsNotFound(err) {
// 			return nil, fmt.Errorf("query record: %w", err)
// 		}
//
// 		rec, err = client.Record.
// 			Create().
// 			SetVideoURL(videoURL).
// 			SetVideoID(videoID).
// 			SetRunID(runID).
// 			SetStatus(record.StatusDownloading).
// 			Save(ctx)
// 		if err != nil {
// 			return nil, fmt.Errorf("create record: %w", err)
// 		}
// 	}
//
// 	return rec, nil
// }
//
// func FetchRecordByVideoID(
// 	ctx context.Context,
// 	client *ent.Client,
// 	videoID string,
// ) (*ent.Record, error) {
// 	rec, err := client.Record.
// 		Query().
// 		Where(record.VideoID(videoID)).
// 		Only(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("fetch record by video ID: %w", err)
// 	}
//
// 	return rec, nil
// }
//
// func DeleteRecordByVideoIDIfDownloadingAndOldRunID(
// 	ctx context.Context,
// 	client *ent.Client,
// 	videoID string,
// 	runID uuid.UUID,
// ) error {
// 	rec, err := FetchRecordByVideoID(ctx, client, videoID)
// 	if err != nil {
// 		if !ent.IsNotFound(err) {
// 			return err
// 		}
// 		return nil
// 	}
// 	if rec.Status != record.StatusDownloading {
// 		return nil
// 	}
// 	if rec.RunID.String() == runID.String() {
// 		return nil
// 	}
// 	if err = client.Record.DeleteOne(rec).Exec(ctx); err != nil {
// 		return fmt.Errorf("failed to delete record: %w", err)
// 	}
//
// 	return nil
// }
