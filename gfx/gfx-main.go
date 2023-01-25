package gfx

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"image"
	_ "image/gif"
	_ "image/png"
	"os"
	"prototyp.com/tomorrows-weather/api"
	"prototyp.com/tomorrows-weather/models"
	"prototyp.com/tomorrows-weather/shaders"
	"time"
)

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
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)
	win.Canvas().SetFragmentShader(shaders.GreyscaleShader)

	delta := 0.0
	last := time.Now()

	sky := generateSky(tomorrowsWeather, currentTimeAtLocation)
	clouds, animationSpeed := generateClouds(tomorrowsWeather)
	drawSky(sky, win)

	for !win.Closed() {

		dt := time.Since(last).Seconds()
		delta += animationSpeed * dt

		last = time.Now()

		drawSky(sky, win)
		//sky.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		drawClouds(win, clouds, delta)
		win.Update()
	}
}

func drawSky(sky *pixel.Sprite, win *pixelgl.Window) {
	sky.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
}

func drawClouds(win *pixelgl.Window, sprites []models.Cloud, delta float64) {
	for _, sprite := range sprites {
		sprite.PositionVec = pixel.V(sprite.Sprite.Picture().Bounds().Center().X+delta+sprite.AnimationDelta, 500)
		sprite.Sprite.Draw(win, pixel.IM.Moved(sprite.PositionVec))
	}
}
