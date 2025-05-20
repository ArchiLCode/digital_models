package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

const migrationsPath = "migrations"

func main() {
	// Загружаем переменные из gateway.env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки gateway.env файла: %v", err)
	}

	// Получаем переменные окружения
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")
	sslmode := os.Getenv("POSTGRES_SSLMODE")

	// Формируем строку подключения
	storagePath := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, dbname, sslmode)

	// Подключаемся к базе данных
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	// Убедимся, что подключение работает
	if err := db.Ping(); err != nil {
		log.Fatalf("Ошибка ping базы данных: %v", err)
	}
	// Выполняем миграции
	if err := goose.Up(db, migrationsPath); err != nil {
		panic(err)
	}

	log.Println("migrated successfully")
}
