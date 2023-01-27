package gfx

import (
	"fmt"

	_ "image/gif"
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
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
	clouds, animationSpeed := generateClouds(tomorrowsWeather)

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		delta += animationSpeed * dt

		last = time.Now()

		if win.Pressed(pixelgl.KeySpace) {
			clouds, animationSpeed = generateClouds(tomorrowsWeather)
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

func drawSky(win *pixelgl.Window, sky *imdraw.IMDraw) {
	sky.Draw(win)
}

func drawClouds(win *pixelgl.Window, clouds *[]models.Cloud, dt float64) {
	ptrClouds := *clouds

	for i, cloud := range ptrClouds {
		halfCloudW := cloud.Sprite.Frame().W() / 2

		if !(cloud.PositionVec.X-(halfCloudW) > win.Bounds().W()) {
			cloud.PositionVec = pixel.V(cloud.PositionVec.X+100*dt, cloud.PositionVec.Y)
			cloud.PositionVec.Scaled(1)
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
