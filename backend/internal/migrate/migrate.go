package main

import (
	"fmt"
	"log"
	"time"

	psq "song/internal/db"
	"song/internal/models"

	"gorm.io/gorm"
)

func main() {
	// Подключение к базе данных
	db, err := psq.ConnectToDB()
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}

	// Создание расширения для UUID (если используется PostgreSQL)
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")

	// Миграция всех таблиц
	err = db.AutoMigrate(
		&models.User{},
		&models.Artist{},
		&models.Album{},
		&models.Song{},
		&models.Playlist{},
		&models.PlaylistSong{},
		&models.UserLikedSong{},
		&models.Genre{},
		&models.SongGenre{},
	)

	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	// Создаем индексы для оптимизации запросов
	db.Exec("CREATE INDEX IF NOT EXISTS idx_songs_title ON songs(title)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_albums_title ON albums(title)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_artists_name ON artists(name)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_songs_play_count ON songs(play_count)")
	db.Exec("CREATE INDEX IF NOT EXISTS idx_playlists_public ON playlists(is_public)")

	// Создание начальных жанров (если нужно)
	createInitialGenres(db)

	// Добавление тестовых данных
	addTestData(db)

	fmt.Println("👍 Миграция успешно завершена")
	fmt.Println("✅ Тестовые данные успешно добавлены")
}

// Создание начальных жанров
func createInitialGenres(db *gorm.DB) {
	genres := []string{
		"Rock", "Pop", "Hip-Hop", "R&B", "Electronic", 
		"Jazz", "Classical", "Country", "Folk", "Metal",
		"Indie", "Blues", "Reggae", "Funk", "Soul",
		"Punk", "Alternative", "Dance", "Disco", "Latin"
	}

	for _, genre := range genres {
		var count int64
		db.Model(&models.Genre{}).Where("name = ?", genre).Count(&count)
		
		if count == 0 {
			db.Create(&models.Genre{
				Name: genre,
			})
		}
	}
}

// Добавление тестовых данных
func addTestData(db *gorm.DB) {
	// Проверяем, есть ли уже данные в базе
	var artistCount int64
	db.Model(&models.Artist{}).Count(&artistCount)
	
	if artistCount > 0 {
		fmt.Println("Тестовые данные уже существуют, пропускаем создание...")
		return
	}

	// Создаем тестовых пользователей
	users := []models.User{
		{
			Username:    "user1",
			Email:       "user1@example.com",
			Password:    "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy", // password: test123
			DisplayName: "Тестовый Пользователь",
		},
		{
			Username:    "user2",
			Email:       "user2@example.com",
			Password:    "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy",
			DisplayName: "Музыкальный Фанат",
		},
	}
	
	for i := range users {
		db.Create(&users[i])
	}

	// Получаем жанры
	var rockGenre, popGenre, hiphopGenre, electronicGenre models.Genre
	db.Where("name = ?", "Rock").First(&rockGenre)
	db.Where("name = ?", "Pop").First(&popGenre)
	db.Where("name = ?", "Hip-Hop").First(&hiphopGenre)
	db.Where("name = ?", "Electronic").First(&electronicGenre)

	// Создаем артистов
	artists := []models.Artist{
		{
			Name:        "Imagine Dragons",
			Description: "Американская инди-рок-группа, образованная в Лас-Вегасе в 2008 году.",
			ImageURL:    "https://example.com/images/imagine_dragons.jpg",
		},
		{
			Name:        "The Weeknd",
			Description: "Канадский певец, автор песен и продюсер.",
			ImageURL:    "https://example.com/images/the_weeknd.jpg",
		},
		{
			Name:        "Billie Eilish",
			Description: "Американская певица и автор песен.",
			ImageURL:    "https://example.com/images/billie_eilish.jpg",
		},
	}
	
	for i := range artists {
		db.Create(&artists[i])
	}

	// Создаем альбомы
	albums := []models.Album{
		{
			Title:       "Night Visions",
			ArtistID:    artists[0].ID,
			ReleaseDate: time.Date(2012, 9, 4, 0, 0, 0, 0, time.UTC),
			CoverURL:    "https://example.com/images/night_visions.jpg",
			AlbumType:   "album",
		},
		{
			Title:       "After Hours",
			ArtistID:    artists[1].ID,
			ReleaseDate: time.Date(2020, 3, 20, 0, 0, 0, 0, time.UTC),
			CoverURL:    "https://example.com/images/after_hours.jpg",
			AlbumType:   "album",
		},
		{
			Title:       "When We All Fall Asleep, Where Do We Go?",
			ArtistID:    artists[2].ID,
			ReleaseDate: time.Date(2019, 3, 29, 0, 0, 0, 0, time.UTC),
			CoverURL:    "https://example.com/images/when_we_all_fall_asleep.jpg",
			AlbumType:   "album",
		},
	}
	
	for i := range albums {
		db.Create(&albums[i])
	}

	// Создаем песни
	songs := []models.Song{
		{
			Title:       "Radioactive",
			ArtistID:    artists[0].ID,
			AlbumID:     albums[0].ID,
			Duration:    188,
			ReleaseDate: albums[0].ReleaseDate,
			AudioURL:    "https://example.com/audio/radioactive.mp3",
			Lyrics:      "I'm waking up to ash and dust\nI wipe my brow and I sweat my rust\nI'm breathing in the chemicals...",
			PlayCount:   1256,
		},
		{
			Title:       "Demons",
			ArtistID:    artists[0].ID,
			AlbumID:     albums[0].ID,
			Duration:    174,
			ReleaseDate: albums[0].ReleaseDate,
			AudioURL:    "https://example.com/audio/demons.mp3",
			Lyrics:      "When the days are cold\nAnd the cards all fold\nAnd the saints we see\nAre all made of gold...",
			PlayCount:   1089,
		},
		{
			Title:       "Blinding Lights",
			ArtistID:    artists[1].ID,
			AlbumID:     albums[1].ID,
			Duration:    203,
			ReleaseDate: time.Date(2019, 11, 29, 0, 0, 0, 0, time.UTC),
			AudioURL:    "https://example.com/audio/blinding_lights.mp3",
			Lyrics:      "I've been tryna call\nI've been on my own for long enough\nMaybe you can show me how to love, maybe...",
			PlayCount:   2345,
		},
		{
			Title:       "Save Your Tears",
			ArtistID:    artists[1].ID,
			AlbumID:     albums[1].ID,
			Duration:    215,
			ReleaseDate: albums[1].ReleaseDate,
			AudioURL:    "https://example.com/audio/save_your_tears.mp3",
			Lyrics:      "I saw you dancing in a crowded room\nYou look so happy when I'm not with you...",
			PlayCount:   1834,
		},
		{
			Title:       "Bad Guy",
			ArtistID:    artists[2].ID,
			AlbumID:     albums[2].ID,
			Duration:    194,
			ReleaseDate: time.Date(2019, 3, 29, 0, 0, 0, 0, time.UTC),
			AudioURL:    "https://example.com/audio/bad_guy.mp3",
			Lyrics:      "White shirt now red, my bloody nose\nSleeping, you're on your tippy toes\nCreeping around like no one knows...",
			PlayCount:   3056,
		},
		{
			Title:       "Bury a Friend",
			ArtistID:    artists[2].ID,
			AlbumID:     albums[2].ID,
			Duration:    193,
			ReleaseDate: time.Date(2019, 1, 30, 0, 0, 0, 0, time.UTC),
			AudioURL:    "https://example.com/audio/bury_a_friend.mp3",
			Lyrics:      "What do you want from me? Why don't you run from me?\nWhat are you wondering? What do you know?...",
			PlayCount:   2198,
		},
	}
	
	for i := range songs {
		db.Create(&songs[i])
	}

	// Связываем песни с жанрами
	songGenres := []models.SongGenre{
		{SongID: songs[0].ID, GenreID: rockGenre.ID},
		{SongID: songs[1].ID, GenreID: rockGenre.ID},
		{SongID: songs[2].ID, GenreID: popGenre.ID},
		{SongID: songs[2].ID, GenreID: electronicGenre.ID},
		{SongID: songs[3].ID, GenreID: popGenre.ID},
		{SongID: songs[4].ID, GenreID: popGenre.ID},
		{SongID: songs[4].ID, GenreID: electronicGenre.ID},
		{SongID: songs[5].ID, GenreID: popGenre.ID},
		{SongID: songs[5].ID, GenreID: electronicGenre.ID},
	}
	
	for i := range songGenres {
		db.Create(&songGenres[i])
	}

	// Создаем плейлисты
	playlists := []models.Playlist{
		{
			Name:        "Мои любимые треки",
			Description: "Коллекция моих любимых песен",
			UserID:      users[0].ID,
			CoverURL:    "https://example.com/images/favorites_playlist.jpg",
			IsPublic:    true,
		},
		{
			Name:        "Для тренировки",
			Description: "Энергичные треки для занятий спортом",
			UserID:      users[0].ID,
			CoverURL:    "https://example.com/images/workout_playlist.jpg",
			IsPublic:    true,
		},
		{
			Name:        "Расслабляющая музыка",
			Description: "Спокойные мелодии для отдыха",
			UserID:      users[1].ID,
			CoverURL:    "https://example.com/images/relax_playlist.jpg",
			IsPublic:    false,
		},
	}
	
	for i := range playlists {
		db.Create(&playlists[i])
	}

	// Добавляем песни в плейлисты
	playlistSongs := []models.PlaylistSong{
		{PlaylistID: playlists[0].ID, SongID: songs[0].ID, Position: 1},
		{PlaylistID: playlists[0].ID, SongID: songs[2].ID, Position: 2},
		{PlaylistID: playlists[0].ID, SongID: songs[4].ID, Position: 3},
		{PlaylistID: playlists[1].ID, SongID: songs[0].ID, Position: 1},
		{PlaylistID: playlists[1].ID, SongID: songs[2].ID, Position: 2},
		{PlaylistID: playlists[2].ID, SongID: songs[1].ID, Position: 1},
		{PlaylistID: playlists[2].ID, SongID: songs[5].ID, Position: 2},
	}
	
	for i := range playlistSongs {
		db.Create(&playlistSongs[i])
	}

	// Добавляем лайки песням
	userLikedSongs := []models.UserLikedSong{
		{UserID: users[0].ID, SongID: songs[0].ID},
		{UserID: users[0].ID, SongID: songs[2].ID},
		{UserID: users[0].ID, SongID: songs[4].ID},
		{UserID: users[1].ID, SongID: songs[1].ID},
		{UserID: users[1].ID, SongID: songs[3].ID},
		{UserID: users[1].ID, SongID: songs[5].ID},
	}
	
	for i := range userLikedSongs {
		db.Create(&userLikedSongs[i])
	}
}