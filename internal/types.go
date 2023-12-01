package internal

type Status string

const (
	StatusDownloading Status = "downloading"
	StatusDownloaded  Status = "downloaded"
	StatusTranscribed Status = "transcribed"
	StatusTranslated  Status = "translated"
	StatusError       Status = "error"
)
