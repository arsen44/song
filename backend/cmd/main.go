package main

import (
	"log"
	psq "song/internal/db"
	"song/internal/handlers"

	"github.com/gin-gonic/gin"
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

func main() {
	// Подключение к базе данных
	db, err := psq.ConnectToDB() // Handle the second return value (error)
	if err != nil {
		// Handle the error properly, e.g., log it and exit
		log.Println("Не удалось подключиться к базе данных: " + err.Error())
	}

	// Заполнение базы данных фиктивными данными
	psq.FillDatabase(db)

	// Инициализация Gin роутера
	r := gin.Default()
	r.Use(CORSMiddleware())
	// Роуты
	r.POST("/add-song/", handlers.AddSong(db))
	r.GET("/songs/", handlers.GetSongs(db))
	r.GET("/songs/:id/text/", handlers.GetSongText(db))
	r.PUT("/songs/:id/", handlers.UpdateSong(db))
	r.DELETE("/songs/:id/", handlers.DeleteSong(db))

	if err := r.Run(":2152"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
