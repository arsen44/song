package repositories

import (
	"encoding/json"
	"errors"
	"net/http"
	"song/internal/models"
)

func FetchSongDetails(group, song string) (models.SongDetail, error) {
	url := "https://external-api.example.com/song?group=" + group + "&song=" + song

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return models.SongDetail{}, errors.New("external API error")
	}
	defer resp.Body.Close()

	var songDetail models.SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return models.SongDetail{}, err
	}

	return songDetail, nil
}
