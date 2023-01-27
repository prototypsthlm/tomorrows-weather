package gfx

import (
	"image"
	_ "image/gif"
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"prototyp.com/tomorrows-weather/api"
	"prototyp.com/tomorrows-weather/models"
)

const WINDOW_SIZE = 468

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

func Run() {
	//load weather
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
	//win.Canvas().SetFragmentShader(shaders.GreyscaleShader)

	delta := 0.0
	last := time.Now()

	sky := generateSky(tomorrowsWeather, currentTimeAtLocation)
	clouds, animationSpeed := generateClouds(tomorrowsWeather)

	drawSky(win, sky)

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		delta += animationSpeed * dt

		last = time.Now()

		drawSky(win, sky)
		drawClouds(win, &clouds, delta)
		removeClouds(win, &clouds)

		win.Update()
	}
}

func drawSky(win *pixelgl.Window, sky *pixel.Sprite) {
	sky.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
}

func removeClouds(win *pixelgl.Window, clouds *[]models.Cloud) {
	ptrClouds := *clouds

	for i, cloud := range *clouds {
		if cloud.PositionX > WINDOW_SIZE {
			cloud.PositionX = 0 - cloud.Sprite.Frame().W()
			ptrClouds[i] = cloud
		}
	}
}

func drawClouds(win *pixelgl.Window, clouds *[]models.Cloud, delta float64) {
	ptrClouds := *clouds

	for i, cloud := range ptrClouds {
		cloud.PositionX += 1
		cloud.PositionVec = pixel.V(cloud.Sprite.Picture().Bounds().Center().X+cloud.AnimationDelta+cloud.PositionX, cloud.PositionY)
		cloud.Sprite.Draw(win, pixel.IM.Moved(cloud.PositionVec))

		ptrClouds[i] = cloud
	}
}
