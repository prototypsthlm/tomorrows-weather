package gfx

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
	"prototyp.com/tomorrows-weather/models"
	"prototyp.com/tomorrows-weather/utils"
)

// assets generated with https://mdigi.tools/gradient-generator/
func generateSky(dailyForecast models.DailyForecast, currentTimeAtLocation int) *imdraw.IMDraw {
	currentTimeAtLocationAsTime := utils.ParseUnixTimestamp(currentTimeAtLocation)
	sunset := utils.ParseUnixTimestamp(dailyForecast.Sunset)
	sunrise := utils.ParseUnixTimestamp(dailyForecast.Sunrise)
	hoursWithSunlight := sunset.Hour() - sunrise.Hour()

	println("hours with sunlight: " + strconv.Itoa(hoursWithSunlight))
	println("current time: " + currentTimeAtLocationAsTime.String())

	//todo: we need to calculate sky based on time of day together with sunrise and sunset
	//todo: cloud density will affect sky color, apply greyscale somehow?
	currentHour := currentTimeAtLocationAsTime.Hour()

	// initialize with default value
	topColor := colornames.Darkblue
	bottomColor := colornames.Lightblue

	if currentHour >= 0 && currentHour < 6 {
		// night
		topColor = colornames.Darkblue
		bottomColor = colornames.Darkblue
	} else if currentHour >= 6 && currentHour < 9 {
		// dawn/sunrise
		topColor = colornames.Darkblue
		bottomColor = colornames.Lightblue
	} else if currentHour >= 9 && currentHour < 16 {
		// midday
		topColor = colornames.Darkblue
		bottomColor = colornames.Lightblue
	} else if currentHour >= 11 && currentHour < 23 {
		// sunset
		topColor = colornames.Lightblue
		bottomColor = colornames.Pink
	}

	imd := imdraw.New(nil)

	// top
	imd.Color = pixel.RGBAModel.Convert(topColor)
	imd.Push(pixel.V(0, WINDOW_SIZE))

	// right
	imd.Color = pixel.RGBAModel.Convert(topColor)
	imd.Push(pixel.V(WINDOW_SIZE, WINDOW_SIZE))

	// bottom
	imd.Color = pixel.RGBAModel.Convert(bottomColor)
	imd.Push(pixel.V(WINDOW_SIZE, 0))

	// left
	imd.Color = pixel.RGBAModel.Convert(bottomColor)
	imd.Push(pixel.V(0, 0))

	imd.Polygon(0)

	return imd
}

func generateClouds(dailyForecast models.DailyForecast) (sprites []models.Cloud, windMultiplier float64) {
	var cloudSprites []models.Cloud
	windMultiplier = calculateAnimationSpeedBasedOnWind(dailyForecast)
	cloudDensity := dailyForecast.Clouds / 10

	//todo: 10 cloud assets is probably overkill, shade background color instead?
	for i := 0; i < cloudDensity; i++ {
		rand.Seed(time.Now().UnixNano())
		//	srcImage := rand.Intn(10-1) + 1
		pic1, err := utils.LoadPicture("./assets/png/test.png")
		//pic1, err := utils.LoadPicture("./assets/png/clouds/" + strconv.Itoa(srcImage) + ".png")
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

	return cloudSprites, windMultiplier
}

func calculateAnimationSpeedBasedOnWind(forecast models.DailyForecast) float64 {
	return forecast.WindSpeed + forecast.WindGust
}
