package repositories

import "song/internal/models"

// SongRepository интерфейс для работы с песнями
type SongRepositoryInterface interface {
	GetAllSongs(albumID uint, songName string, page, limit int) ([]models.Song, error)
	DeleteSongs(uint) error
}
