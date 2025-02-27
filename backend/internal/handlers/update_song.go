package handlers

import (
	"net/http"
	"song/internal/models"
	"song/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Обновить песню
// @Description Обновляет информацию о песне по ID
// @Tags models.Song
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Param song body models.Song true "Обновленные данные песни"
// @Success 200 {object} models.Song
// @Failure 400 {string} string "Ошибка в запросе"
// @Failure 404 {string} string "Песня не найдена"
// @Router /songs/{id}/ [put]
func UpdateSong(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var input models.SongUpdate
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := repositories.UpdateSong(db, id, input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update song"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Song updated successfully!"})
	}
}
