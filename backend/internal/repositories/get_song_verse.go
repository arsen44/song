package repositories

import (
	"errors"
	"song/internal/models"
	"strings"

	"gorm.io/gorm"
)

// GetSongVerse извлекает конкретный куплет песни
func GetSongVerse(db *gorm.DB, id string, verse int) (string, error) {
	var song models.Song
	if err := db.First(&song, "id = ?", id).Error; err != nil {
		return "", err
	}

	// Разделение текста на куплеты
	verses := strings.Split(song.Text, "\n\n")
	if verse <= 0 || verse > len(verses) {
		return "", errors.New("verse out of range")
	}

	return verses[verse-1], nil
}
