package models

import (
	"time"

	"gorm.io/gorm"
)

// User представляет пользователя сервиса
type User struct {
	gorm.Model
	Username    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Email       string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	DisplayName string `gorm:"type:varchar(100)"`
	Playlists   []Playlist
	LikedSongs  []Song `gorm:"many2many:user_liked_songs;"`
}

// Artist представляет исполнителя
type Artist struct {
	gorm.Model
	Name        string `gorm:"type:varchar(150);not null"`
	Description string `gorm:"type:text"`
	ImageURL    string `gorm:"type:varchar(255)"`
	Albums      []Album
	Songs       []Song
}

// Album представляет альбом
type Album struct {
	gorm.Model
	Title       string    `gorm:"type:varchar(150);not null"`
	ArtistID    uint      `gorm:"index;not null"`
	Artist      Artist    `gorm:"foreignKey:ArtistID"`
	ReleaseDate time.Time `gorm:"index;not null"`
	CoverURL    string    `gorm:"type:varchar(255)"`
	AlbumType   string    `gorm:"type:varchar(50);not null"` // album, single, EP
	Songs       []Song    `gorm:"foreignKey:AlbumID"`
}

// Song представляет песню
type Song struct {
	gorm.Model
	Title       string     `gorm:"type:varchar(150);not null" json:"title"`
	ArtistID    uint       `gorm:"index;not null" json:"artist_id"`
	Artist      Artist     `gorm:"foreignKey:ArtistID" json:"-"`
	AlbumID     uint       `gorm:"index" json:"album_id"`
	Album       Album      `gorm:"foreignKey:AlbumID" json:"-"`
	Duration    int        `gorm:"not null" json:"duration"` // длительность в секундах
	ReleaseDate time.Time  `gorm:"index;not null" json:"release_date"`
	AudioURL    string     `gorm:"type:varchar(255);not null" json:"audio_url"`
	Lyrics      string     `gorm:"type:text" json:"lyrics"`
	PlayCount   int        `gorm:"default:0" json:"play_count"`
	Genres      []Genre    `gorm:"many2many:song_genres;"`
	Playlists   []Playlist `gorm:"many2many:playlist_songs;"`
}

// Playlist представляет плейлист
type Playlist struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(150);not null"`
	Description string    `gorm:"type:text"`
	UserID      uint      `gorm:"index;not null"`
	User        User      `gorm:"foreignKey:UserID"`
	CoverURL    string    `gorm:"type:varchar(255)"`
	IsPublic    bool      `gorm:"default:true"`
	Songs       []Song    `gorm:"many2many:playlist_songs;"`
	CreatedAt   time.Time `gorm:"index"`
	UpdatedAt   time.Time
}

// PlaylistSong связывает песни с плейлистами с порядковым номером
type PlaylistSong struct {
	PlaylistID uint     `gorm:"primaryKey"`
	SongID     uint     `gorm:"primaryKey"`
	Position   int      `gorm:"not null"`
	Playlist   Playlist `gorm:"foreignKey:PlaylistID"`
	Song       Song     `gorm:"foreignKey:SongID"`
	CreatedAt  time.Time
}

// UserLikedSong связывает пользователей с понравившимися песнями
type UserLikedSong struct {
	UserID  uint      `gorm:"primaryKey"`
	SongID  uint      `gorm:"primaryKey"`
	User    User      `gorm:"foreignKey:UserID"`
	Song    Song      `gorm:"foreignKey:SongID"`
	LikedAt time.Time `gorm:"not null;autoCreateTime"`
}

// SongGenre связывает песни с жанрами
type SongGenre struct {
	SongID  uint  `gorm:"primaryKey"`
	GenreID uint  `gorm:"primaryKey"`
	Song    Song  `gorm:"foreignKey:SongID"`
	Genre   Genre `gorm:"foreignKey:GenreID"`
}

// Genre представляет музыкальный жанр
type Genre struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Songs []Song `gorm:"many2many:song_genres;"`
}

// SongUpdate используется для обновления полей Song
type SongUpdate struct {
	Title       *string    `json:"title,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Lyrics      *string    `json:"lyrics,omitempty"`
	AudioURL    *string    `json:"audio_url,omitempty"`
	Duration    *int       `json:"duration,omitempty"`
}

// SongDetail представляет детальную информацию о песне
type SongDetail struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Artist      string    `json:"artist"`
	Album       string    `json:"album"`
	ReleaseDate time.Time `json:"release_date"`
	Duration    int       `json:"duration"`
	Lyrics      string    `json:"lyrics"`
	AudioURL    string    `json:"audio_url"`
	CoverURL    string    `json:"cover_url"`
}

// SongInput представляет входные данные для добавления песни
type SongInput struct {
	Title    string `json:"title" binding:"required"`
	ArtistID uint   `json:"artist_id" binding:"required"`
	AlbumID  uint   `json:"album_id"`
	Duration int    `json:"duration" binding:"required"`
	AudioURL string `json:"audio_url" binding:"required"`
	Lyrics   string `json:"lyrics"`
	GenreIDs []uint `json:"genre_ids"`
}
