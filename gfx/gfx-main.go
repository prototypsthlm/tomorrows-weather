package gfx

import (
	"fmt"
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

	pic1, err := loadPicture("./assets/clouds/clouds-1.png")
	if err != nil {
		panic(err)
	}

	pic2, err := loadPicture("./assets/clouds/clouds-2.png")
	if err != nil {
		panic(err)
	}

	sprite := pixel.NewSprite(pic1, pic1.Bounds())
	sprite2 := pixel.NewSprite(pic2, pic2.Bounds())

	delta := 0.0
	delta2 := 0.0
	last := time.Now()

	win.Clear(colornames.Lightgray)
	sprite.Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {

		dt := time.Since(last).Seconds()
		delta += 18 * dt
		delta2 += 36 * dt
		last = time.Now()

		drawClouds(delta, delta2, win, sprite, sprite2)

		win.Update()
	}
}

func respawnClouds(pixel.Vec) {

}

func drawClouds(delta float64, delta2 float64, win *pixelgl.Window, sprite *pixel.Sprite, sprite2 *pixel.Sprite) {
	win.Clear(colornames.Lightslategray)

	v := pixel.V(sprite.Picture().Bounds().Center().X+delta, 500)

	winX := win.Bounds().Max.X

	fmt.Println("winX: ", winX)

	if v.X > winX {
		v2 := pixel.V(-500, -1.5)

		v = pixel.V(v2.X+delta2, v2.Y)

		sprite.Draw(win, pixel.IM.Moved(v))
		fmt.Printf("%v vector", v)
	} else {
		sprite.Draw(win, pixel.IM.Moved(v))
	}
}
