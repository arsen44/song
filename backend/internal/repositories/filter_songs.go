package repositories

import (
	"song/internal/models"

	"gorm.io/gorm"
)

func FilterSongs(db *gorm.DB, albumID uint, song string, page, limit int) ([]models.Song, error) {
	var songs []models.Song
	query := db.Model(&models.Song{}).Preload("Album")

	if albumID != 0 {
		query = query.Where("album_id = ?", albumID)
	}
	if song != "" {
		query = query.Where("song_title ILIKE ?", "%"+song+"%")
	}

	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}
