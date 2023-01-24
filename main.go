package main

import (
	"github.com/joho/godotenv"
	"prototyp.com/tomorrows-weather/api"
)

func main() {
	godotenv.Load(".env")
	api.GetTomorrowsWeather()
}
