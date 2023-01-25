package gfx

import (
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"image/color"
	"math/rand"
	"prototyp.com/tomorrows-weather/models"
	"prototyp.com/tomorrows-weather/utils"
	"strconv"
	"time"
)

func generateBackground(dailyForecast models.DailyForecast) color.RGBA {
	//todo: return color by time of day
	timeOfForecast := utils.ParseUnixTimestamp(dailyForecast.Dt) //todo: might not be needed? Go with curren time at lat lon instead?
	println("forecast time: " + timeOfForecast.String())
	return colornames.Mediumpurple
}

func generateClouds(dailyForecast models.DailyForecast) (sprites []models.Cloud, animationSpeed float64) {
	var cloudSprites []models.Cloud
	animationSpeed = 10
	cloudDensity := dailyForecast.Clouds / 10

	//todo: 10 cloud assets is probably overkill, shade background color instead?
	for i := 0; i < cloudDensity; i++ {
		rand.Seed(time.Now().UnixNano())
		srcImage := rand.Intn(10-1) + 1
		pic1, err := loadPicture("./assets/png/clouds/" + strconv.Itoa(srcImage) + ".png")
		if err != nil {
			panic(err)
		}

		animationDelta := float64(rand.Intn(10-1) + 1)
		//todo: rand select position via bounds?

		var cloud = models.Cloud{
			Sprite:         pixel.NewSprite(pic1, pic1.Bounds()),
			AnimationDelta: animationDelta,
		}
		cloudSprites = append(cloudSprites, cloud)
	}

	//todo: calc animationSpeed based on wind
	//todo: should be slightly random to create layers

	return cloudSprites, animationSpeed
}
