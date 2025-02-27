package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	POSTGRES_HOST     = os.Getenv("POSTGRES_HOST")
	POSTGRES_USER     = os.Getenv("POSTGRES_USER")
	POSTGRES_PORT     = os.Getenv("DB_PORT")
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	POSTGRES_NAME     = os.Getenv("POSTGRES_NAME")
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

	log.Println("Миграция схемы завершена успешно.")
	return db, nil
}
