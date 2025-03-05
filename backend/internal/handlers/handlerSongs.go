package handlers

import (
	"net/http"
	"song/internal/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SongHandler обработчик запросов связанных с песнями
type SongHandlers struct {
	songRepo repositories.SongRepositoryInterface
}

func NewSongHandlers(songRepo repositories.SongRepositoryInterface) SongsHandlersInterface {
	return &SongHandlers{songRepo: songRepo}
}

// @Summary Получить список песен
// @Description Возвращает список всех песен
// @Tags songs
// @Produce json
// @Success 200 {array} models.Song
// @Router /songs/ [get]
func (h *SongHandlers) GetAllSongs(c *gin.Context) {
	albumIDStr := c.Query("album_id")
	var albumID uint
	if albumIDStr != "" {
		parsedID, err := strconv.ParseUint(albumIDStr, 10, 32)
		if err == nil {
			albumID = uint(parsedID)
		}
	}

	song := c.Query("song")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	songs, err := h.songRepo.GetAllSongs(albumID, song, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve songs"})
		return
	}

	c.JSON(http.StatusOK, songs)
}

func (h *SongHandlers) DeleteSongs(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID"})
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid song ID format"})
		return
	}

	if err := h.songRepo.DeleteSongs(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Song deleted successfully"})
}
