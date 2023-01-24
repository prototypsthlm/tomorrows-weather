package gfx

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/png"
	"os"
	"prototyp.com/tomorrows-weather/api"
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

	pic1, err := loadPicture("./assets/svg/cloud1.png")
	if err != nil {
		panic(err)
	}

	/*
		pic2, err := loadPicture("./assets/svg/cloud2.png")
		if err != nil {
			panic(err)
		}

	*/

	sprite := pixel.NewSprite(pic1, pic1.Bounds())
	//sprite2 := pixel.NewSprite(pic2, pic2.Bounds())

	delta := 0.0
	delta2 := 0.0
	last := time.Now()

	skyColor := generateBackground(tomorrowsWeather)
	clouds, animationSpeed := generateClouds(tomorrowsWeather)
	win.Clear(skyColor)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	//sprite2.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {

		dt := time.Since(last).Seconds()
		delta += animationSpeed * dt
		delta2 += animationSpeed * dt
		last = time.Now()

		//drawTwoClouds(delta, delta2, win, sprite, sprite2, skyColor)
		win.Clear(skyColor)
		drawClouds(win, clouds, skyColor)
		//mat = mat.Moved(win.Bounds().Center())
		//sprite.Draw(win, mat)
		win.Update()
	}
}

func drawClouds(win *pixelgl.Window, sprites []*pixel.Sprite, skyColor color.RGBA) {
	for _, sprite := range sprites {
		//todo: to animate, extend with updated delta-X
		v := pixel.V(sprite.Picture().Bounds().Center().X, 500)
		sprite.Draw(win, pixel.IM.Moved(v))
	}
}

func drawTwoClouds(delta float64, delta2 float64, win *pixelgl.Window, sprite *pixel.Sprite, sprite2 *pixel.Sprite, skyColor color.RGBA) {
	win.Clear(skyColor)

	//mat := pixel.IM
	//mat = mat.Rotated(pixel.ZV, delta)
	v := pixel.V(sprite.Picture().Bounds().Center().X+delta, 500)
	v2 := pixel.V(sprite.Picture().Bounds().Center().X+delta2, 300)
	sprite.Draw(win, pixel.IM.Moved(v))
	sprite2.Draw(win, pixel.IM.Moved(v2))
}
