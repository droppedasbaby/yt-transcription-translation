package internal

import (
	"context"
	"fmt"
	"sync"

	"github.com/sashabaranov/go-openai"
	"go.uber.org/zap"
)

type Download struct {
	VideoID              string
	VideoURL             string
	Transcript           string
	RequestedTranslation string
	Translation          string
	Path                 string
	Status               Status
}

type DownloadManager struct {
	ctx    context.Context
	mu     sync.Mutex
	client *openai.Client
	logger *zap.Logger
	store  map[string](*Download)
}

func NewDownloadManager(parentLogger *zap.Logger, openaiKey string, ctx context.Context) *DownloadManager {
	// TODO: Make testible
	client := openai.NewClient(openaiKey)

	logger := parentLogger.With(zap.String("component", "internal.NewDownloadManager"))
	logger.Info("Initializing download manager")

	return &DownloadManager{
		mu:     sync.Mutex{},
		store:  make(map[string]*Download),
		logger: logger,
		client: client,
		ctx:    ctx,
	}
}

func (dm *DownloadManager) AddDownload(vURL string, translation string) (*Download, error) {
	dm.logger.Info("Adding download to download manager", zap.String("video_url", vURL), zap.String("translation", translation))
	vID, err := GetVideoID(vURL)
	if err != nil {
		return nil, fmt.Errorf("AddDownload failed: %w", err)
	}

	d := &Download{
		VideoID:              vID,
		VideoURL:             vURL,
		RequestedTranslation: translation,
		Status:               StatusDownloading,
	}

	go dm.manageDownload(d)
	return d, nil
}

func (dm *DownloadManager) GetResults(vID string) (*Download, error) {
	dm.logger.Info("Getting results", zap.String("video_id", vID))
	dm.mu.Lock()
	d, exists := dm.store[vID]
	if !exists {
		dm.mu.Unlock()
		return nil, fmt.Errorf("GetResults failed: download does not exist")
	}
	dm.mu.Unlock()
	copy := *d
	return &copy, nil
}

func (dm *DownloadManager) manageDownload(d *Download) {
	dm.logger.Info("Managing download", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))

	key := generateKey(d.VideoID, d.RequestedTranslation)

	dm.mu.Lock()
	od, exists := dm.store[key]
	if exists && od.Status == StatusError {
		dm.logger.Warn("Download already exists and is in error state. Deleting and re-downloading", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))
		delete(dm.store, key)
	}
	dm.mu.Unlock()

	if !exists {
		dm.mu.Lock()
		dm.logger.Info("Adding download to store", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))
		dm.store[key] = d
		dm.mu.Unlock()

		dm.logger.Info("Downloading video", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))
		path, _, err := DownloadVideo(d.VideoID)

		dm.mu.Lock()
		if err != nil {
			dm.logger.Error("Error downloading video", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation), zap.Error(err))
			dm.store[key].Status = StatusError
			return
		}
		dm.logger.Info("Successfully downloaded video", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))
		dm.store[key].Path = path
		dm.store[key].Status = StatusDownloaded
		dm.mu.Unlock()
	}

	if dm.store[key].Status == StatusDownloading {
		dm.logger.Info("Getting transcript", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))
		transcript, err := dm.getTranscript(d.Path)
		if err != nil {
			dm.setStatusToError(key)
			return
		}

		dm.mu.Lock()
		dm.store[key].Status = StatusTranscribed
		dm.store[key].Transcript = transcript
		dm.mu.Unlock()
	}
	if dm.store[key].Status == StatusTranscribed {
		if dm.store[key].RequestedTranslation != "" {
			dm.logger.Info("Getting translation", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))
			translation, err := dm.getTranslation(d.Transcript, d.RequestedTranslation)
			if err != nil {
				dm.setStatusToError(key)
				return
			}
			dm.mu.Lock()
			dm.store[key].Translation = translation
			dm.mu.Unlock()
		} else {
			dm.logger.Info("No translation requested", zap.String("video_id", d.VideoID), zap.String("translation", d.RequestedTranslation))
		}

		dm.mu.Lock()
		dm.store[key].Status = StatusTranslated
		dm.mu.Unlock()
	}
}

func (dm *DownloadManager) getTranscript(path string) (string, error) {
	req := openai.AudioRequest{
		Model:    openai.Whisper1,
		FilePath: path,
		Format:   openai.AudioResponseFormatSRT,
	}
	resp, err := dm.client.CreateTranscription(dm.ctx, req)
	if err != nil {
		return "", fmt.Errorf("getTranscript failed: %w", err)
	}
	return resp.Text, nil
}

func (dm *DownloadManager) getTranslation(transcript string, translation string) (string, error) {
	prompt := "Translate everything in ```...``` to " + translation + "\n."
	prompt += "Keep the same format, SRT."
	prompt += "Return nothing without the results of the translation,"
	prompt += " here is what I want translate:\n"
	prompt += "```\n"
	prompt += transcript
	prompt += "\n```"

	req := openai.CompletionRequest{
		Model:  openai.GPT432K0613,
		Prompt: prompt,
	}
	resp, err := dm.client.CreateCompletion(dm.ctx, req)
	if err != nil {
		return "", fmt.Errorf("getTranslation failed: %w", err)
	}
	return resp.Choices[0].Text, nil
}

func (dm *DownloadManager) setStatusToError(key string) {
	dm.mu.Lock()
	dm.store[key].Status = StatusError
	dm.mu.Unlock()
}

func generateKey(vID string, translation string) string {
	return vID + "+" + translation
}
