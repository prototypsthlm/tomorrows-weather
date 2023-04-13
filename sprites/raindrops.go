package sprites

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/sandnuggah/tomorrows-weather/config"
	"github.com/sandnuggah/tomorrows-weather/models"
	"github.com/sandnuggah/tomorrows-weather/utils"
)

const (
	RAIN_SPEED = 1
)

type Raindrop struct {
	inited bool
	X      float64
	Y      float64
	VelX   float64
	VelY   float64
	Alpha  int
}

func (raindrop *Raindrop) init() {
	defer func() {
		raindrop.inited = true
	}()
	raindrop.Y = utils.Randf(
		-config.WindowHeight,
		-20,
	)
	raindrop.X = utils.Randf(0, config.WindowHeight)
	raindrop.VelX = utils.Randf(-0.5, 0.5)
	raindrop.VelY = utils.Randf(15, 20)
	raindrop.Alpha = utils.Rand(5, 50)
}

func (raindrop *Raindrop) Update(forecast models.DailyForecast) {
	if !raindrop.inited {
		raindrop.init()
	}
	raindrop.X = raindrop.X + raindrop.VelX
	raindrop.Y = raindrop.Y + (raindrop.VelY * RAIN_SPEED)
	if config.WindowHeight < raindrop.Y {
		raindrop.init()
	}
}

func (raindrop *Raindrop) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(
		screen,
		raindrop.X,
		raindrop.Y,
		raindrop.X,
		raindrop.Y+utils.Randf(10, 20),
		color.RGBA{
			R: uint8(200),
			G: uint8(200),
			B: uint8(200),
			A: uint8(raindrop.Alpha),
		},
	)
}

type Raindrops struct {
	Raindrops []*Raindrop
	Num       int
}

func (raindrops *Raindrops) Update(forecast models.DailyForecast) {
	for i := 0; i < raindrops.Num; i++ {
		raindrops.Raindrops[i].Update(forecast)
	}
}

func (raindrops *Raindrops) Draw(screen *ebiten.Image) {
	for i := 0; i < raindrops.Num; i++ {
		raindrops.Raindrops[i].Draw(screen)
	}
}
