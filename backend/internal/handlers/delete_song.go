package handlers

import (
	"net/http"
	"song/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Удалить песню
// @Description Удаляет песню по ID
// @Tags songs
// @Accept json
// @Produce json
// @Param id path int true "ID песни"
// @Success 200 {string} string "Песня успешно удалена"
// @Failure 404 {string} string "Песня не найдена"
// @Router /songs/{id}/ [delete]
func DeleteSong(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := repositories.DeleteSong(db, id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete song"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully!"})
	}
}
