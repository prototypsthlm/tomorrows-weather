package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"prototyp.com/tomorrows-weather/models"
)

func GetTomorrowsWeather() models.DailyForecast {
	//Timișoara  lat and long
	var lat = "45.7489"
	var lon = "21.2087"
	var apiKey = os.Getenv("OPEN_WEATHER_API_KEY")
	apiURL := "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lon + "&exclude=minutely,hourly&units=metric&appid=" + apiKey
	print(apiURL)
	response, err := http.Get(apiURL)

	if err != nil {
		fmt.Print(err.Error())
		return models.DailyForecast{}
	}

	var result models.Forecast
	json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return models.DailyForecast{}
	}
	return result.Daily[1]
}
