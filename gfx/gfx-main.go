package gfx

import (
	"fmt"

	_ "image/gif"
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-co-op/gocron"
	"prototyp.com/tomorrows-weather/api"
	"prototyp.com/tomorrows-weather/models"
)

const WINDOW_SIZE = 468

var tomorrowsWeather models.DailyForecast
var currentTimeAtLocation int

func setCurrentWeatherBasedOnForecast() {
	tomorrowsWeather, currentTimeAtLocation = api.GetTomorrowsWeather()
}

func Run() {
	setCurrentWeatherBasedOnForecast() //init weather
	cfg := pixelgl.WindowConfig{
		Title:  "Tomorrows Weather",
		Bounds: pixel.R(0, 0, WINDOW_SIZE, WINDOW_SIZE),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)

	delta := 0.0
	last := time.Now()

	UpdateWeatherOnInterval() //start scheduled job to update weather
	sky := generateSky(tomorrowsWeather, currentTimeAtLocation)
	clouds, windMultiplier := generateClouds(tomorrowsWeather)

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		delta += windMultiplier * dt

		last = time.Now()

		if win.Pressed(pixelgl.KeySpace) {
			clouds, windMultiplier = generateClouds(tomorrowsWeather)
		}

		drawSky(win, sky)
		drawClouds(win, &clouds, dt)

		win.Update()
	}
}

func UpdateWeatherOnInterval() {
	updateFrequencyInMinutes := 10
	fmt.Println("Starting cron job to update weather every", updateFrequencyInMinutes, "minutes")
	s := gocron.NewScheduler(time.UTC)

	s.Every(updateFrequencyInMinutes).Minutes().Do(func() {
		println("updating weather")
		setCurrentWeatherBasedOnForecast()
	})
	s.StartAsync()
}

func drawSky(win *pixelgl.Window, sky *pixel.Sprite) {
	sky.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
}

func drawClouds(win *pixelgl.Window, clouds *[]models.Cloud, dt float64) {
	ptrClouds := *clouds

	for i, cloud := range ptrClouds {
		halfCloudW := cloud.Sprite.Frame().W() / 2

		if !(cloud.PositionVec.X-(halfCloudW) > win.Bounds().W()) {
			newXPosition := cloud.PositionVec.X + 100*dt
			cloud.PositionVec = pixel.V(newXPosition, cloud.PositionVec.Y)
			cloud.Sprite.Draw(win, pixel.IM.Moved(cloud.PositionVec))
		} else {
			cloud.PositionVec = pixel.V(-halfCloudW, 100)
		}

		ptrClouds[i] = cloud
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func drawRain(win *pixelgl.Window) {
	panic("not implemented")
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func drawSnow(win *pixelgl.Window) {
	panic("not implemented")
}
