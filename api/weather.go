package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"prototyp.com/tomorrows-weather/models"
)

func GetWeatherReport() models.DailyForecast {
	//Bucharest lat and long
	var lat = "44.4268"
	var lon = "26.1025"
	var apiKey = os.Getenv("OPEN_WEATHER_API_KEY")
	apiURL := "https://api.openweathermap.org/data/3.0/onecall?lat=" + lat + "&lon=" + lon + "&exclude=minutely,hourly&units=metric&appid=" + apiKey
	print(apiURL)
	response, err := http.Get(apiURL)

	if err != nil {
		fmt.Print(err.Error())
	}

	var result models.DailyForecast
	json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Current)
	return result
}
