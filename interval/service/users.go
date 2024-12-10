package service

import (
	"database/sql"
	"log"
	"net/http"

	models "project/interval/models"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context, db *sql.DB) {
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var usernames []string
	for rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			log.Fatal(err)
		}
		usernames = append(usernames, username)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	var result models.ResultUsers
	result.Users = usernames

	c.JSON(http.StatusOK, result)
}
