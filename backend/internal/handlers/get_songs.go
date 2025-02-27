package handlers

import (
	"net/http"
	"song/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Получить список песен
// @Description Возвращает список всех песен
// @Tags songs
// @Produce json
// @Success 200 {array} models.Song
// @Router /songs/ [get]
func GetSongs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		group := c.Query("group")
		song := c.Query("song")
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

		songs, err := repositories.FilterSongs(db, group, song, page, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
			return
		}

		c.JSON(http.StatusOK, songs)
	}
}
