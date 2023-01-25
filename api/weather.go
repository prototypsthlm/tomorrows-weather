package api

import (
	"encoding/json"
	"net/http"
	"os"
	"prototyp.com/tomorrows-weather/models"
)

func GetTomorrowsWeather() (tomorrowsWeather models.DailyForecast, currentTime int) {
	//Timișoara  lat and long
	var lat = "45.7489"
	var lon = "21.2087"
	var apiKey = os.Getenv("OPEN_WEATHER_API_KEY")
	apiURL := "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lon + "&exclude=minutely,hourly&units=metric&appid=" + apiKey
	response, err := http.Get(apiURL)

	if err != nil {
		println("err: " + err.Error())
		return models.DailyForecast{}, 0
	}

	var result models.Forecast
	json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		println("err: " + err.Error())
		return models.DailyForecast{}, 0
	}
	return result.Daily[1], result.Current.Dt
}
