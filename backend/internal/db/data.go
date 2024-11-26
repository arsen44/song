package db

import (
	"log"
	"song/internal/models"
	"song/internal/utils"
	"time"

	"gorm.io/gorm"
)

// mustParseReleaseDate обрабатывает ошибки парсинга или завершает выполнение
func mustParseReleaseDate(dateStr string) time.Time {
	parsedDate, err := utils.ParseReleaseDate(dateStr)
	if err != nil {
		log.Fatalf("Ошибка парсинга даты %s: %v", dateStr, err)
	}
	return parsedDate
}

// Фиктивные данные для заполнения базы
var songs = []models.Song{
	{
		GroupName:   "Muse",
		SongName:    "Supermassive Black Hole",
		ReleaseDate: mustParseReleaseDate("16.07.2006"),
		Text:        "Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\n...",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	},
	{
		GroupName:   "Queen",
		SongName:    "Bohemian Rhapsody",
		ReleaseDate: mustParseReleaseDate("31.10.1975"),
		Text:        "Is this the real life? Is this just fantasy?\nCaught in a landslide, no escape from reality...",
		Link:        "https://www.youtube.com/watch?v=fJ9rUzIMcZQ",
	},
	{
		GroupName:   "The Beatles",
		SongName:    "Hey Jude",
		ReleaseDate: mustParseReleaseDate("26.08.1968"),
		Text:        "Hey Jude, don't make it bad\nTake a sad song and make it better...",
		Link:        "https://www.youtube.com/watch?v=A_MjCqQoLLA",
	},
}

// FillDatabase заполняет базу данных фиктивными данными
func FillDatabase(db *gorm.DB) {
	for _, song := range songs {
		var existingSong models.Song
		err := db.Where("group_name = ? AND song_name = ?", song.GroupName, song.SongName).First(&existingSong).Error
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(&song).Error; err != nil {
				log.Printf("Ошибка при добавлении песни %s - %s: %v", song.GroupName, song.SongName, err)
			} else {
				log.Printf("Добавлена песня: %s - %s", song.GroupName, song.SongName)
			}
		}
	}
}
