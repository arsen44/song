package repositories

import (
	"song/internal/models"

	"gorm.io/gorm"
)

// songRepository реализация интерфейса SongRepository
type SongRepository struct {
	db *gorm.DB
}

// NewSongRepository создает новый экземпляр репозитория песен
func NewSongRepository(db *gorm.DB) SongRepositoryInterface {
	return &SongRepository{db: db}
}

// FilterSongs фильтрует песни по заданным параметрам
// Метод реализует интерфейс SongRepository
func (r *SongRepository) GetAllSongs(albumID uint, song string, page, limit int) ([]models.Song, error) {
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

// DeleteSong удаляет песню по её ID
func (r *SongRepository) DeleteSongs(id uint) error {
	tx := r.db.Begin()

	// Удаляем связанные записи в таблице user_liked_songs
	if err := tx.Where("song_id = ?", id).Delete(&models.UserLikedSong{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Удаляем связанные записи в таблице playlist_songs
	if err := tx.Where("song_id = ?", id).Delete(&models.PlaylistSong{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Удаляем связанные записи в таблице song_genres
	if err := tx.Where("song_id = ?", id).Delete(&models.SongGenre{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Удаляем песню
	if err := tx.Unscoped().Delete(&models.Song{}, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
