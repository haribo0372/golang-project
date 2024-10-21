package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Загрузка переменных окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	dbuser := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	appId := os.Getenv("APP_ID")

	// Подключение к базе данных PostgreSQL
	db, err := setupDatabase(dbuser, dbpassword, dbname)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	// Создаем экземпляр Gin
	r := gin.Default()

	// Маршруты для регистрации и авторизации
	r.POST("/register", func(c *gin.Context) {
		register(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		login(c, db)
	})

	authorized := r.Group("/")
	authorized.Use(authMiddleware())
	{
		authorized.GET("/weather", func(c *gin.Context) {
			weather(c, appId)
		})
	}

	// Запуск сервера на порту 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
