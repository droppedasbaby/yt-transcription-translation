package types

import "github.com/GrewalAS/yt-transcription-translation/internal"

type RecordPostBody struct {
	URL              string `json:"url"`
	OriginalLanguage string `json:"original_language"`
}
type RecordResponse struct{}

type RecordResultsQueryParams struct {
	VideoID string `json:"video_id"`
}

type RecordResultsResponse struct {
	Status internal.Status `json:"status"`
}

type ProcessVideoIDWorkerResults struct{}
