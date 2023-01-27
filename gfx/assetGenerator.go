package gfx

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"prototyp.com/tomorrows-weather/models"
	"prototyp.com/tomorrows-weather/utils"
)

// assets generated with https://mdigi.tools/gradient-generator/
func generateSky(dailyForecast models.DailyForecast, currentTimeAtLocation int) *pixel.Sprite {
	currentTimeAtLocationAsTime := utils.ParseUnixTimestamp(currentTimeAtLocation)
	sunset := utils.ParseUnixTimestamp(dailyForecast.Sunset)
	sunrise := utils.ParseUnixTimestamp(dailyForecast.Sunrise)
	hoursWithSunlight := sunset.Hour() - sunrise.Hour()
	println("hours with sunlight: " + strconv.Itoa(hoursWithSunlight))
	println("current time: " + currentTimeAtLocationAsTime.String())

	//todo: we need to calculate sky based on time of day together with sunrise and sunset
	//todo: cloud density will affect sky color, apply greyscale somehow?
	currentHour := currentTimeAtLocationAsTime.Hour()
	if currentHour >= 0 && currentHour < 6 {
		skyPic, _ := utils.LoadPicture("./assets/png/sky/night.png")
		return pixel.NewSprite(skyPic, skyPic.Bounds())
		//night

	} else if currentHour >= 6 && currentHour < 9 {
		//dawn/sunrise
		skyPic, _ := utils.LoadPicture("./assets/png/sky/sunrise.png")
		return pixel.NewSprite(skyPic, skyPic.Bounds())

	} else if currentHour >= 9 && currentHour < 16 {
		//midday
		skyPic, _ := utils.LoadPicture("./assets/png/sky/midday.png")
		return pixel.NewSprite(skyPic, skyPic.Bounds())
	} else {
		//sunset
		skyPic, _ := utils.LoadPicture("./assets/png/sky/sunset.png")
		return pixel.NewSprite(skyPic, skyPic.Bounds())
	}

	//default
	skyPic, _ := utils.LoadPicture("./assets/png/sky/midday.png")
	return pixel.NewSprite(skyPic, skyPic.Bounds())
}

func generateClouds(dailyForecast models.DailyForecast) (sprites []models.Cloud, animationSpeed float64) {
	var cloudSprites []models.Cloud
	animationSpeed = calculateAnimationSpeedBasedOnWind(dailyForecast)
	cloudDensity := dailyForecast.Clouds / 10

	//todo: 10 cloud assets is probably overkill, shade background color instead?
	for i := 0; i < cloudDensity; i++ {
		rand.Seed(time.Now().UnixNano())
		srcImage := rand.Intn(10-1) + 1
		pic1, err := utils.LoadPicture("./assets/png/clouds/" + strconv.Itoa(srcImage) + ".png")
		if err != nil {
			panic(err)
		}
		rand.Seed(time.Now().UnixNano())
		animationDelta := float64(rand.Intn(10-1) + 1)
		rand.Seed(time.Now().UnixNano())
		PositionX := float64(rand.Intn(468))
		rand.Seed(time.Now().UnixNano())
		PositionY := float64(rand.Intn(468))

		var cloud = models.Cloud{
			Sprite:         pixel.NewSprite(pic1, pic1.Bounds()),
			AnimationDelta: animationDelta,
			PositionVec:    pixel.V(PositionX, PositionY),
		}
		cloudSprites = append(cloudSprites, cloud)
	}

	return cloudSprites, animationSpeed
}

func calculateAnimationSpeedBasedOnWind(forecast models.DailyForecast) float64 {
	return forecast.WindSpeed + forecast.WindGust
}
