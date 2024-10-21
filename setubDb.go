package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Функция для настройки подключения к PostgreSQL
func setupDatabase(dbuser string, dbpassword string, dbname string) (*sql.DB, error) {

	if dbuser == "" || dbpassword == "" || dbname == "" {
		return nil, fmt.Errorf("не все необходимые переменные окружения установлены")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbuser, dbpassword, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Успешное подключение к базе данных!")
	return db, nil
}
