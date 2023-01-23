package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"prototyp.com/tomorrows-weather/api"
)

func main() {
	godotenv.Load(".env")
	fmt.Println("Hello World")
	api.GetWeatherReport()
}
