package models

type Weather struct {
	WeatherDescription string `json:"description"`
}

type Results struct {
	City       string  `json:"name"`
	Results    Main    `json:"main"`
	Visibility float64 `json:"visibility"`
	Wind       Wind    `json:"wind"`
	WindSpeed  string
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Main struct {
	Temperature          float64 `json:"temperature"`
	MinTemperature       float64 `json:"temp_min"`
	MaxTemperature       float64 `json:"temp_max"`
	FeelsLikeTemperature float64 `json:"feels_like"`
}
