package handlers

import (
	"net/http"
	"song/internal/models"
	"song/internal/repositories"
	"song/internal/services"
	"song/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Добавить песню
// @Description Добавляет новую песню в базу данных
// @Tags songs
// @Param song body models.Song  true "Данные песни"
// @Success 200 {object} models.Song
// @Failure 400 {string} string "Ошибка в запросе"
// @Router /add-song/ [post]
func AddSong(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.SongInput

		// Парсинг JSON-запроса
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Вызов внешнего API для получения деталей песни
		songDetail, err := services.FetchSongDetails(input.Group, input.Song)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch song details: " + err.Error()})
			return
		}

		releaseDate, err := utils.ParseReleaseDate(songDetail.ReleaseDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid release date format: " + err.Error()})
			return
		}

		// Преобразование songDetail в модель Song
		song := models.Song{
			GroupName:   input.Group,
			SongName:    input.Song,
			ReleaseDate: releaseDate,
			Text:        songDetail.Text,
			Link:        songDetail.Link,
		}

		// Сохранение данных в базу через репозиторий
		if err := repositories.SaveSong(db, song); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save song to database: " + err.Error()})
			return
		}

		// Успешный ответ
		c.JSON(http.StatusCreated, gin.H{"message": "Song added successfully!"})
	}
}
