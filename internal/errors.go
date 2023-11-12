package internal

import (
	"fmt"
)

type GetVideoIDError struct {
	URL string
	Msg string
}

func (e GetVideoIDError) Error() string {
	return fmt.Sprintf("%s: %s", e.Msg, e.URL)
}

func ErrInvalidYouTubeURL(url string) error {
	return GetVideoIDError{
		URL: url,
		Msg: "Not a valid YouTube URL",
	}
}

func ErrVideoIDNotFound(url string) error {
	return GetVideoIDError{
		URL: url,
		Msg: "Video ID not found",
	}
}
