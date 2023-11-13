package server

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/GrewalAS/yt-transcription-translation/internal"
	"github.com/GrewalAS/yt-transcription-translation/types"
)

func (s *Server) startHandler(w http.ResponseWriter, r *http.Request) {
	var body types.RecordPostBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		s.logger.Error("Failed to decode request body", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	videoID, err := internal.GetVideoID(body.URL)
	if err != nil {
		s.logger.Error("Failed to get video ID", zap.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = DeleteRecordByVideoIDIfDownloadingAndOldRunID(r.Context(), s.client, videoID, s.runID); err != nil {
		s.logger.Error("Failed to delete record", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rec, err := FetchOrCreateRecordByVideoID(r.Context(), s.client, body.URL, videoID, s.runID)
	if err != nil {
		s.logger.Error("Failed to fetch or create record", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := types.RecordResponse{
		VideoURL:     rec.VideoURL,
		VideoID:      rec.VideoID,
		FileLocation: rec.FileLocation,
		Status:       rec.Status,
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(resp); err != nil {
		s.logger.Error("Failed to encode response", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
