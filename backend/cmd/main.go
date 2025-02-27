package main

import (
	"log"
	_ "song/docs"
	psq "song/internal/db"
	"song/internal/handlers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Song API
// @version 1.0
// @description API для управления песнями
// @host localhost:2152
// @BasePath /

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

	// Инициализация Gin роутера
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/songs/", handlers.GetSongs(db))

	if err := r.Run(":2152"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
