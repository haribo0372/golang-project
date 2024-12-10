package setup_db

import (
	"database/sql"
	"fmt"

	config "project/config"

	_ "github.com/lib/pq"
)

func SetupDatabase(envVariables config.EnvModel) (*sql.DB, error) {

	if envVariables.DbName == "" || envVariables.DbPassword == "" || envVariables.DbUser == "" {
		return nil, fmt.Errorf("не все необходимые переменные окружения установлены")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", envVariables.DbUser, envVariables.DbPassword, envVariables.DbName)
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
