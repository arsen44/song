package repositories

import (
	"song/internal/models"

	"gorm.io/gorm"
)

// SongRepository интерфейс для работы с песнями
type SongRepository interface {
	FilterSongs(albumID uint, songName string, page, limit int) ([]models.Song, error)
}

// songRepository реализация интерфейса SongRepository
type songRepository struct {
	db *gorm.DB
}

// NewSongRepository создает новый экземпляр репозитория песен
func NewSongRepository(db *gorm.DB) SongRepository {
	return &songRepository{db: db}
}

// FilterSongs фильтрует песни по заданным параметрам
// Метод реализует интерфейс SongRepository
func (r *songRepository) FilterSongs(albumID uint, song string, page, limit int) ([]models.Song, error) {
	var songs []models.Song
	query := r.db.Model(&models.Song{}).Preload("Album")

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
