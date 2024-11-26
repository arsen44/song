package db

import (
	"fmt"
	"log"
	"song/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var (
// 	POSTGRES_HOST     = os.Getenv("POSTGRES_HOST")
// 	POSTGRES_USER     = os.Getenv("POSTGRES_USER")
// 	POSTGRES_PORT     = os.Getenv("DB_PORT")
// 	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
// 	POSTGRES_NAME     = os.Getenv("POSTGRES_NAME")
// )

// Жестко закодированные значения переменных окружения
var (
	POSTGRES_HOST     = "192.168.1.67"
	POSTGRES_USER     = "postgres"
	POSTGRES_PORT     = "5432"
	POSTGRES_PASSWORD = "170888"
	POSTGRES_NAME     = "postgres"
)

func ConnectToDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_NAME,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
		return nil, err
	}

	// Автоматическое создание или обновление таблиц
	err = db.AutoMigrate(&models.Song{}) // Добавьте здесь все свои модели
	if err != nil {
		log.Fatalf("Ошибка при миграции схемы базы данных: %v", err)
		return nil, err
	}

	log.Println("Миграция схемы завершена успешно.")
	return db, nil
}
