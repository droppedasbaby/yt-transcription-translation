package server

import (
	"net/http"
)

func (s *Server) startHandler(w http.ResponseWriter, r *http.Request) {
	// var body types.RecordPostBody
	// if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
	// 	s.logger.Error("Failed to decode request body", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	//
	// videoID, err := internal.GetVideoID(body.URL)
	// if err != nil {
	// 	s.logger.Error("Failed to get video ID", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	//
	// if err = DeleteRecordByVideoIDIfDownloadingAndOldRunID(r.Context(), s.client, videoID, s.runID); err != nil {
	// 	s.logger.Error("Failed to delete record", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	// rec, err := FetchOrCreateRecordByVideoID(r.Context(), s.client, body.URL, videoID, s.runID)
	// if err != nil {
	// 	s.logger.Error("Failed to fetch or create record", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	// go internal.ProcessVideoID(s.ctx, s.logger, rec)
	//
	// resp := types.RecordResponse{
	// 	VideoURL:     rec.VideoURL,
	// 	VideoID:      rec.VideoID,
	// 	FileLocation: rec.FileLocation,
	// 	Status:       rec.Status,
	// }
	//
	// w.WriteHeader(http.StatusOK)
	// if err = json.NewEncoder(w).Encode(resp); err != nil {
	// 	s.logger.Error("Failed to encode response", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func (s *Server) resultsHandler(w http.ResponseWriter, r *http.Request) {
	// var queryParams types.RecordResultsQueryParams
	// if err := json.NewDecoder(r.Body).Decode(&queryParams); err != nil {
	// 	s.logger.Error("Failed to decode request body", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }
	//
	// rec, err := FetchRecordByVideoID(r.Context(), s.client, queryParams.VideoID)
	// if err != nil {
	// 	s.logger.Error("Failed to fetch record", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//
	// resp := types.RecordResultsResponse{
	// 	Status: rec.Status,
	// }
	//
	// w.WriteHeader(http.StatusOK)
	// if err = json.NewEncoder(w).Encode(resp); err != nil {
	// 	s.logger.Error("Failed to encode response", zap.Error(err))
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}
