package internal

import (
	"context"
	"fmt"
	"sync"
)

type Download struct {
	VideoID              string
	VideoURL             string
	Transcript           string
	RequestedTranslation string
	Translation          string
	Path                 string
	Status               Status
	Error                error
}

type DownloadManager struct {
	ctx   context.Context
	mu    sync.RWMutex
	store map[string](*Download)
}

func NewDownloadManager() *DownloadManager {
	return &DownloadManager{
		// TODO: Add logger and logging
		mu:    sync.RWMutex{},
		store: make(map[string]*Download),
	}
}

func (dm *DownloadManager) AddDownload(vURL string, translation string) (*Download, error) {
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

func (dm *DownloadManager) GetResults(vID string) (*Download, bool) {
	return nil, false
}

func (dm *DownloadManager) manageDownload(d *Download) {
	// TODO: Use lock properly
	// TODO: Make more efficient, no need for multiple downloads
	key := generateKey(d.VideoID, d.RequestedTranslation)

	od, exists := dm.store[key]
	if exists && od.Status == StatusError {
		delete(dm.store, key)
	}

	if !exists {
		dm.store[key] = d

		path, _, err := DownloadVideo(d.VideoID)
		if err != nil {
			dm.store[key].Status = StatusError
			dm.store[key].Error = err
		}

		dm.store[key].Status = StatusDownloaded
	}

	if d.Status == StatusDownloaded {
		// TODO: Add translation
	}

	if d.Status == StatusTranscribed {
		// TODO: Add translation
	}

	if d.Status == StatusTranslated {
		// TODO: Anything needs to be done here?
	}
}

func generateKey(vID string, translation string) string {
	return vID + "+" + translation
}
