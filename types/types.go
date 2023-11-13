package types

import "github.com/GrewalAS/yt-transcription-translation/ent/record"

type RecordPostBody struct {
	URL string `json:"url"`
}

type RecordResponse struct {
	VideoURL     string        `json:"video_url"`
	VideoID      string        `json:"video_id"`
	FileLocation string        `json:"file_location"`
	Status       record.Status `json:"status,omitempty"`
}
