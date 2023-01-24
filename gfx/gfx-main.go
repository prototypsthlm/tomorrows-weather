package gfx

import (
	"image"
	_ "image/gif"
	_ "image/png"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

	pic2, err := loadPicture("./assets/svg/cloud2.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic1, pic1.Bounds())
	sprite2 := pixel.NewSprite(pic2, pic2.Bounds())

	delta := 0.0
	delta2 := 0.0
	last := time.Now()

	win.Clear(colornames.Skyblue)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	//sprite2.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {

		dt := time.Since(last).Seconds()
		delta += 18 * dt
		delta2 += 36 * dt
		last = time.Now()

		drawClouds(delta, delta2, win, sprite, sprite2)
		//mat = mat.Moved(win.Bounds().Center())
		//sprite.Draw(win, mat)
		win.Update()
	}
}

func drawClouds(delta float64, delta2 float64, win *pixelgl.Window, sprite *pixel.Sprite, sprite2 *pixel.Sprite) {
	win.Clear(colornames.Mediumpurple)

	//mat := pixel.IM
	//mat = mat.Rotated(pixel.ZV, delta)
	v := pixel.V(sprite.Picture().Bounds().Center().X+delta, 500)
	v2 := pixel.V(sprite.Picture().Bounds().Center().X+delta2, 300)
	sprite.Draw(win, pixel.IM.Moved(v))
	sprite2.Draw(win, pixel.IM.Moved(v2))
}
