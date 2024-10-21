package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	_ "github.com/lib/pq"
)

func weather(c *gin.Context, appid string) {
	sCity := c.Query("s_city")
	if sCity == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Параметр 's_city' обязателен"})
		return
	}

	// Выполняем запрос к OpenWeatherMap API
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
		fmt.Println(err.Error())
		return
	}

	if resp.StatusCode() != http.StatusOK {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Город не найден или ошибка запроса"})
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки ответа от OpenWeather API"})
		return
	}

	// Извлекаем данные из ответа API
	weatherInfo := gin.H{
		"city":                   sCity,
		"weather_description":    result["weather"].([]interface{})[0].(map[string]interface{})["description"].(string),
		"temperature":            fmt.Sprintf("%.2f C°", result["main"].(map[string]interface{})["temp"].(float64)),
		"min_temperature":        fmt.Sprintf("%.2f C°", result["main"].(map[string]interface{})["temp_min"].(float64)),
		"max_temperature":        fmt.Sprintf("%.2f C°", result["main"].(map[string]interface{})["temp_max"].(float64)),
		"feels_like_temperature": fmt.Sprintf("%.2f C°", result["main"].(map[string]interface{})["feels_like"].(float64)),
		"visibility":             fmt.Sprintf("%v м", result["visibility"]),
		"wind_speed":             fmt.Sprintf("%.2f м/с", result["wind"].(map[string]interface{})["speed"].(float64)),
	}

	if gust, ok := result["wind"].(map[string]interface{})["gust"]; ok {
		weatherInfo["wind_gust"] = fmt.Sprintf("%.2f м/с", gust.(float64))
	} else {
		weatherInfo["wind_gust"] = "Нет данных"
	}

	// Возвращаем данные клиенту
	c.JSON(http.StatusOK, weatherInfo)
}
