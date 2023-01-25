package gfx

import (
	"image"
	_ "image/gif"
	_ "image/png"
	"os"
	"prototyp.com/tomorrows-weather/api"
	"prototyp.com/tomorrows-weather/models"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	tomorrowsWeather := api.GetTomorrowsWeather()

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

	delta := 0.0
	last := time.Now()

	skyColor := generateBackground(tomorrowsWeather)
	clouds, animationSpeed := generateClouds(tomorrowsWeather)
	win.Clear(skyColor)

	for !win.Closed() {

		dt := time.Since(last).Seconds()
		delta += animationSpeed * dt

		last = time.Now()

		win.Clear(skyColor)
		drawClouds(win, clouds, delta)
		win.Update()
	}
}

func drawClouds(win *pixelgl.Window, sprites []models.Cloud, delta float64) {
	for _, sprite := range sprites {
		sprite.PositionVec = pixel.V(sprite.Sprite.Picture().Bounds().Center().X+delta+sprite.AnimationDelta, 500)
		sprite.Sprite.Draw(win, pixel.IM.Moved(sprite.PositionVec))
	}
}
