package types

type RecordStatus string

const (
	RecordStatusDownloading RecordStatus = "downloading"
	RecordStatusDownloaded  RecordStatus = "downloaded"
)

type Record struct {
	VideoURL     string       `json:"video_url"`
	VideoID      string       `json:"video_id"`
	FileLocation string       `json:"file_location"`
	Status       RecordStatus `json:"status,omitempty"`
}
