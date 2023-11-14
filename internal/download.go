package internal

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/kkdai/youtube/v2"
)

func GetVideoID(vURL string) (string, error) {
	parsedURL, err := url.Parse(vURL)
	if err != nil {
		return "", fmt.Errorf("parse URL: %w", err)
	}

	domain := parsedURL.Hostname()
	if domain != "youtube.com" && domain != "www.youtube.com" && domain != "youtu.be" {
		return "", ErrInvalidYouTubeURL(vURL)
	}

	var id string
	if domain == "youtu.be" {
		id = parsedURL.Path[1:]
	} else {
		queryParams := parsedURL.Query()
		id = queryParams.Get("v")
	}

	if id == "" {
		return "", ErrVideoIDNotFound(vURL)
	}

	return id, nil
}

func createFile(id string) (*os.File, error) {
	fullPath, err := CreateDirAndGetFullPath(fmt.Sprintf("%s.mp4", id))
	if err != nil {
		return nil, err
	}

	if err = os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return nil, fmt.Errorf("create directory: %w", err)
	}
	if _, err = os.Stat(fullPath); err == nil {
		if err = os.Remove(fullPath); err != nil {
			return nil, fmt.Errorf("delete existing file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return nil, fmt.Errorf("check file existence: %w", err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return nil, fmt.Errorf("create file: %w", err)
	}

	return file, nil
}

func DownloadVideo(ID string) (string, int64, error) {
	client := youtube.Client{}
	video, err := client.GetVideo(ID)
	if err != nil {
		return "", -1, fmt.Errorf("fetch video info: %w", err)
	}

	formats := video.Formats.WithAudioChannels()
	stream, length, err := client.GetStream(video, &formats[0])
	if err != nil {
		return "", length, fmt.Errorf("fetch video stream: %w", err)
	}
	defer stream.Close()

	file, err := createFile(ID)
	if err != nil {
		return "", length, fmt.Errorf("create file: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		return "", length, fmt.Errorf("copy stream to file: %w", err)
	}

	return file.Name(), length, nil
}

func TranscribeVideo() error {
	return nil
}

