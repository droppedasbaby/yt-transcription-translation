package server

import (
	"encoding/json"
	"net/http"

	"github.com/GrewalAS/yt-transcription-translation/ent"
)

func startHandler(w http.ResponseWriter, r *http.Request) {
	var newRec ent.Record
	if err := json.NewDecoder(r.Body).Decode(&newRec); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
