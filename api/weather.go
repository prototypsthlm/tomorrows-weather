package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetWeatherReport() {
	//bukarest lat and long
	var lat = "44.4268"
	var lon = "26.1025"
	var apiKey = os.Getenv("OPEN_WEATHER_API_KEY")
	apiURL := "https://api.openweathermap.org/data/2.5/weather?lat=" + lat + "&lon=" + lon + "&appid=" + apiKey
	print(apiURL)
	response, err := http.Get(apiURL)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
}
