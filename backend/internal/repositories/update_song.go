package repositories

import (
	"song/internal/models"

	"gorm.io/gorm"
)

// UpdateSong обновляет песню в базе данных по ID
func UpdateSong(db *gorm.DB, id string, input models.SongUpdate) error {
	// Поиск существующей записи
	var song models.Song
	if err := db.First(&song, "id = ?", id).Error; err != nil {
		return err
	}

	// Обновление полей
	if input.GroupName != nil {
		song.GroupName = *input.GroupName
	}
	if input.SongName != nil {
		song.SongName = *input.SongName
	}
	if input.ReleaseDate != nil {
		song.ReleaseDate = *input.ReleaseDate
	}
	if input.Text != nil {
		song.Text = *input.Text
	}
	if input.Link != nil {
		song.Link = *input.Link
	}

	// Сохранение изменений
	if err := db.Save(&song).Error; err != nil {
		return err
	}

	return nil
}
