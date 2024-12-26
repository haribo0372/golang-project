package main

import (
	"log"

	config "project/config"

	setupDb "project/interval/db"
	handlers "project/interval/handlers"

	_ "github.com/lib/pq"
)

func main() {
	envVariables := config.LoadEnv()

	db, err := setupDb.SetupDatabase(envVariables)

	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	r := handlers.SetupHandlers(db, envVariables)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
