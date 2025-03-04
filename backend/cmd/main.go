package main

import (
	"log"
	_ "song/docs"
	psq "song/internal/db"

	"song/internal/routes"
)

// @title Song API
// @version 1.0
// @description API для управления песнями
// @host localhost:2152
// @BasePath /

func main() {
	// Подключение к базе данных
	dbConn, err := psq.ConnectToDB() // Handle the second return value (error)
	if err != nil {
		// Handle the error properly, e.g., log it and exit
		log.Println("Не удалось подключиться к базе данных: " + err.Error())
	}

	// Создаем маршрутизатор с передачей соединения с базой данных
	router := routes.SetupRouter(dbConn)

	if err := router.Run(":2152"); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
