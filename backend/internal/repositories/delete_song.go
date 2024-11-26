package repositories

import (
	"log"
	"song/internal/models"

	"gorm.io/gorm"
)

func DeleteSong(db *gorm.DB, id string) error {
	log.Printf("Attempting to delete song with ID: %s", id) // Логируем ID
	result := db.Delete(&models.Song{}, "id = ?", id)
	if result.Error != nil {
		log.Printf("Error deleting song: %v", result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		log.Printf("No rows affected, song with ID %s not found", id)
		return gorm.ErrRecordNotFound
	}
	log.Printf("Deleted song with ID: %s", id)
	return nil
}
