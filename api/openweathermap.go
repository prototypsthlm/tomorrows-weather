package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sandnuggah/tomorrows-weather/config"
	"github.com/sandnuggah/tomorrows-weather/models"
)

func GetForecast(lon, lat float64) (forecast models.DailyForecast, timezone string) {
	fmt.Printf(
		"openweathermap/GetForecast lon=%10f lat=%10f\n",
		lon,
		lat,
	)
	client := http.Client{}
	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/3.0/onecall?lat=%f&lon=%f&exclude=minutely,hourly&units=metric&appid=%s",
		lat,
		lon,
		config.OpenWeatherMapApiKey,
	)
	response, err := client.Get(url)
	if err != nil {
		log.Println(err)
		return models.DailyForecast{}, "Europe/Stockholm"
	}
	var result models.Forecast
	json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		log.Println(err)
		return models.DailyForecast{}, "Europe/Stockholm"
	}
	return result.Daily[1], result.Timezone
}
