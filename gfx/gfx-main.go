package gfx

import (
	_ "image/gif"
	_ "image/png"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"prototyp.com/tomorrows-weather/api"
	"prototyp.com/tomorrows-weather/models"
)

const WINDOW_SIZE = 468

func Run() {
	tomorrowsWeather, currentTimeAtLocation := api.GetTomorrowsWeather()

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

func drawSky(win *pixelgl.Window, sky *pixel.Sprite) {
	sky.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
}

func drawClouds(win *pixelgl.Window, clouds *[]models.Cloud, dt float64) {
	ptrClouds := *clouds

	for i, cloud := range ptrClouds {
		halfCloudW := cloud.Sprite.Frame().W() / 2

		if !(cloud.PositionVec.X-(halfCloudW) > win.Bounds().W()) {
			cloud.PositionVec = pixel.V(cloud.PositionVec.X+100*dt, cloud.PositionVec.Y)
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
