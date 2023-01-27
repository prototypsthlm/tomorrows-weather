package gfx

import (
	"fmt"
	"math/rand"

	_ "image/gif"
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/go-co-op/gocron"
	"golang.org/x/image/colornames"
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
		drawRain(win, dt)

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
			newPositionX := cloud.PositionVec.X + 100*dt
			cloud.PositionVec = pixel.V(newPositionX, cloud.PositionVec.Y)

			mat := pixel.IM
			mat = mat.Scaled(pixel.V(newPositionX, cloud.PositionVec.Y), cloud.ScaleFactor)
			mat = mat.Moved(cloud.PositionVec)

			cloud.Sprite.Draw(win, mat)
		} else {
			rand.Seed(time.Now().UnixNano())
			newPositionY := float64(rand.Intn(468))
			cloud.PositionVec = pixel.V(-halfCloudW, newPositionY)
		}

		ptrClouds[i] = cloud
	}
}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func drawRain(win *pixelgl.Window, dt float64) {
	imd := imdraw.New(nil)

	for i := 1; i <= 10; i++ {
		imd.Color = pixel.RGBAModel.Convert(colornames.White)
		imd.Push(pixel.V(float64(i*10), 10))

		imd.Color = pixel.RGBAModel.Convert(pixel.Alpha(0))
		imd.Push(pixel.V(float64(i*10), 38))

		imd.Polygon(2)

		imd.Draw(win)
	}

}

//lint:ignore U1000 Ignore unused function temporarily for debugging
func drawSnow(win *pixelgl.Window) {
	panic("not implemented")
}
