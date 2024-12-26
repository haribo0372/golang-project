package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "project/config"
	models "project/interval/models"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/lib/pq"
)

func GetWeather(c *gin.Context, envVariables config.EnvModel) {
	logger := log.Default()

	sCity := c.Query("s_city")
	if sCity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 's_city' обязателен"})
		return
	}

	appid := envVariables.AppId

	client := resty.New()
	apiURL := "http://api.openweathermap.org/data/2.5/weather"
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"q":     sCity,
			"units": "metric",
			"lang":  "ru",
			"APPID": appid,
		}).
		Get(apiURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка запроса к OpenWeather API"})
		logger.Fatal("Запрос к openweathermap был выполнен", err.Error())
		return
	}

	if resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Город не найден или ошибка запроса"})
		return
	}

	var result models.Results
	err = json.Unmarshal([]byte(resp.Body()), &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки ответа от OpenWeather API"})
		logger.Fatalf(err.Error())
		return
	}

	logger.Println(resp)

	if result.Wind.Speed != 0 {
		result.WindSpeed = fmt.Sprintf("%.2f м/с", result.Wind.Speed)
	} else {
		result.WindSpeed = "Нет данных"
	}

	c.JSON(http.StatusOK, result)
}
