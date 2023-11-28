package internal_test

import (
	"errors"
	"github.com/GrewalAS/yt-transcription-translation/internal"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetVideoID_Success(t *testing.T) {
	tests := []struct {
		name string
		url  string
		want string
	}{
		{
			name: "Valid YouTube URL",
			url:  "http://youtube.com/watch?v=TestID",
			want: "TestID",
		},
		{
			name: "YouTube URL with additional parameters",
			url:  "http://youtube.com/watch?v=TestID&feature=youtu.be",
			want: "TestID",
		},
		{
			name: "YouTube URL with short format",
			url:  "http://youtu.be/TestID",
			want: "TestID",
		},
		{
			name: "YouTube URL with www",
			url:  "http://www.youtube.com/watch?v=TestID",
			want: "TestID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := internal.GetVideoID(tt.url)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetVideoID_Failure(t *testing.T) {
	tests := []struct {
		name          string
		url           string
		expectedError string
	}{
		{
			name:          "Invalid YouTube URL",
			url:           "http://invalid.url",
			expectedError: "Not a valid YouTube URL: http://invalid.url",
		},
		{
			name:          "Video ID not found",
			url:           "http://youtube.com/watch?v=",
			expectedError: "Video ID not found: http://youtube.com/watch?v=",
		},
		{
			name:          "URL without http protocol",
			url:           "youtube.com/watch?v=TestID",
			expectedError: "Not a valid YouTube URL: youtube.com/watch?v=TestID",
		},
		{
			name:          "Empty string",
			url:           "",
			expectedError: "Not a valid YouTube URL: ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := internal.GetVideoID(tt.url)
			if err == nil {
				t.Fatalf("Expected error but got nil")
			}

			var getVideoIDErr internal.GetVideoIDError
			if !errors.As(err, &getVideoIDErr) {
				t.Fatalf("Expected error of type GetVideoIDError, got: %T", err)
			}

			assert.Equal(t, tt.expectedError, err.Error())
			assert.Equal(t, tt.url, getVideoIDErr.URL)
		})
	}
}

func TestDownloadVideo(t *testing.T) {
	id := "dQw4w9WgXcQ" // Rick Astley's "Never Gonna Give You Up"

	filename, length, err := internal.DownloadVideo(id)

	assert.NoError(t, err)
	assert.NotEqual(t, -1, length)
	assert.NotEmpty(t, filename)

	info, err := os.Stat(filename)
	assert.NoError(t, err)
	assert.NotZero(t, info.Size())

	os.Remove(filename)
}
