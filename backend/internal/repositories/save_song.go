package repositories

import (
	"song/internal/models"

	"gorm.io/gorm"
)

// SaveSong сохраняет песню в базу данных
func SaveSong(db *gorm.DB, song models.Song) error {
	return db.Create(&song).Error
}
