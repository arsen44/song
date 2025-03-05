package routes

import (
	"song/internal/handlers"
	"song/internal/repositories"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT,DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SetupRouter настраивает и возвращает маршрутизатор приложения
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	// Инициализация Gin роутера
	r.Use(CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Инициализация репозиториев
	songRepo := repositories.NewSongRepository(db)

	// Инициализация обработчиков
	songHandler := handlers.NewSongHandlers(songRepo)

	// Настройка маршрутов
	r.GET("/songs/", songHandler.GetAllSongs)
	r.DELETE("/songs/:id", songHandler.DeleteSongs)
	return r
}
