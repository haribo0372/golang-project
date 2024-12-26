package handlers

import (
	"database/sql"

	config "project/config"
	auth "project/interval/auth"
	service "project/interval/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func SetupHandlers(db *sql.DB, envVariables config.EnvModel) *gin.Engine {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		auth.RegisterProcess(c, db)
	})
	r.POST("/login", func(c *gin.Context) {
		auth.LoginProcess(c, db)
	})

	authorized := r.Group("/")
	authorized.Use(auth.AuthMiddleware())
	{
		authorized.GET("/weather", func(c *gin.Context) {
			service.GetWeather(c, envVariables)
		})

		r.GET("/users", func(c *gin.Context) {
			service.GetAllUsers(c, db)
		})
	}

	return r
}
