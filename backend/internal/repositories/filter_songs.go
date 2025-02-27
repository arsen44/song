package repositories

import (
	"song/internal/models"

	"gorm.io/gorm"
)

func FilterSongs(db *gorm.DB, group, song string, page, limit int) ([]models.Song, error) {
	var songs []models.Song
	query := db.Model(&models.Song{})

	if group != "" {
		query = query.Where("group_name ILIKE ?", "%"+group+"%")
	}
	if song != "" {
		query = query.Where("song_name ILIKE ?", "%"+song+"%")
	}

	offset := (page - 1) * limit
	if err := query.Limit(limit).Offset(offset).Find(&songs).Error; err != nil {
		return nil, err
	}
	return songs, nil
}
