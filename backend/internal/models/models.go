package models

import "time"

type Song struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	GroupName   string    `gorm:"type:text;not null" json:"group_name"`
	SongName    string    `gorm:"type:text;not null" json:"song_name"`
	ReleaseDate time.Time `gorm:"type:date;not null" json:"release_date"`
	Text        string    `gorm:"type:text;not null" json:"text"`
	Link        string    `gorm:"type:text;not null" json:"link"`
}

// SongUpdate используется для обновления полей Song
type SongUpdate struct {
	GroupName   *string    `json:"group_name,omitempty"`
	SongName    *string    `json:"song_name,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Text        *string    `json:"text,omitempty"`
	Link        *string    `json:"link,omitempty"`
}

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"` // Дата релиза
	Text        string `json:"text"`        // Текст песни
	Link        string `json:"link"`        // Ссылка на источник
}

// SongInput представляет входные данные для добавления песни
type SongInput struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}
