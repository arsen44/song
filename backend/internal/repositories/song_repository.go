package repositories

import (
	"encoding/json"
	"errors"
	"net/http"
	"song/internal/models"
	"strings"

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

// DeleteSong удаляет песню из базы данных по ID
func DeleteSong(db *gorm.DB, id string) error {
	// Удаление песни по ID
	if err := db.Delete(&models.Song{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

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

// SaveSong сохраняет песню в базу данных
func SaveSong(db *gorm.DB, song models.Song) error {
	return db.Create(&song).Error
}

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
