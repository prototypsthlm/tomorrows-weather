package main

import (
	"github.com/faiface/pixel/pixelgl"
	"github.com/joho/godotenv"
	"prototyp.com/tomorrows-weather/gfx"
)

func main() {
	godotenv.Load(".env")
	//api.GetTomorrowsWeather()
	pixelgl.Run(gfx.Run)

}
