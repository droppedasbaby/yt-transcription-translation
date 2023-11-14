package types

import "github.com/GrewalAS/yt-transcription-translation/ent/record"

type RecordPostBody struct {
	URL              string `json:"url"`
	OriginalLanguage string `json:"original_language"`
}

type RecordResponse struct {
	VideoURL                  string        `json:"video_url"`
	VideoID                   string        `json:"video_id"`
	FileLocation              string        `json:"file_location"`
	Status                    record.Status `json:"status,omitempty"`
	TranslationTargetLanguage string        `json:"translation_target_language,omitempty"`
}

type RecordResultsQueryParams struct {
	VideoID string `json:"video_id"`
}

type RecordResultsResponse struct {
	Status record.Status `json:"status"`
}

type ProcessVideoIDWorkerResults struct {
	VideoID      string
	FileLocation string
	Transcript   string
	Translation  string
}
