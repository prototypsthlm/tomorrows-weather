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

	pic, err := loadPicture("./assets/svg/clouds.gif")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic, pic.Bounds())
	sprite2 := pixel.NewSprite(pic, pic.Bounds())

	delta := 0.0
	delta2 := 0.0
	last := time.Now()

	win.Clear(colornames.Skyblue)
	//sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		delta += 18 * dt
		delta2 += 36 * dt

		win.Clear(colornames.Mediumpurple)

		//mat := pixel.IM
		//mat = mat.Rotated(pixel.ZV, delta)
		v := pixel.V(sprite.Picture().Bounds().Center().X+delta, 500)
		v2 := pixel.V(sprite.Picture().Bounds().Center().X+delta2, 300)
		sprite.Draw(win, pixel.IM.Moved(v))
		sprite2.Draw(win, pixel.IM.Moved(v2))
		//mat = mat.Moved(win.Bounds().Center())
		//sprite.Draw(win, mat)

		win.Update()
	}
}
