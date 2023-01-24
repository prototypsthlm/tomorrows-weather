package gfx

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"prototyp.com/tomorrows-weather/models"
	"prototyp.com/tomorrows-weather/utils"
)

func generateBackground(dailyForecast models.DailyForecast) color.RGBA {
	//todo: return color by time of day
	timeOfForecast := utils.ParseUnixTimestamp(dailyForecast.Dt) //todo: might not be needed? Go with curren time at lat lon instead?
	println("forecast time: " + timeOfForecast.String())
	return colornames.Mediumpurple
}

func generateClouds(dailyForecast models.DailyForecast) (sprites []*pixel.Sprite, animationSpeed float64) {
	var cloudSprites []*pixel.Sprite
	animationSpeed = 0
	println("dailyForecast.Clouds" + string(dailyForecast.Clouds))
	clouds := dailyForecast.Clouds / 10

	//cloud density
	for i := 0; i < clouds; i++ {
		//todo: rand select image
		//todo: rand select position
		pic1, err := loadPicture("./assets/svg/cloud1.png")
		if err != nil {
			panic(err)
		}

		cloudSprites = append(cloudSprites, pixel.NewSprite(pic1, pic1.Bounds()))
	}

	return cloudSprites, animationSpeed
}
