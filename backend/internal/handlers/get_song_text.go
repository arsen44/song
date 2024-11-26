package handlers

import (
	"net/http"
	"song/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetSongText(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		verse, _ := strconv.Atoi(c.DefaultQuery("verse", "1"))

		text, err := repositories.GetSongVerse(db, id, verse)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve song text"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"verse": text})
	}
}
