package sprites

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sandnuggah/tomorrows-weather/config"
	"github.com/sandnuggah/tomorrows-weather/models"
	"github.com/sandnuggah/tomorrows-weather/utils"
)

type Cloud struct {
	inited  bool
	ImgW    int
	ImgH    int
	PosX    float64
	PosY    float64
	VelX    float64
	VelY    float64
	Scale   float64
	Opacity float64
	Texture *ebiten.Image
	op      ebiten.DrawImageOptions
}

func (cloud *Cloud) init() {
	defer func() {
		cloud.inited = true
	}()

	cloud.Scale = utils.Randf(0.6, 0.9)

	cloud.PosY = utils.Randf(
		-200,
		config.WindowHeight+200,
	)

	cloud.op.ColorM.Scale(
		1,
		1,
		1,
		cloud.Opacity,
	)
	cloud.op.ColorM.ChangeHSV(0, 0.5, 0.9)
}

func (cloud *Cloud) Update(forecast models.DailyForecast) {
	if !cloud.inited {
		cloud.init()
	}
	if cloud.PosX > config.WindowWidth/cloud.Scale+2000 {
		cloud.init()
		cloud.PosX = -float64(cloud.ImgW)
	}
	if cloud.PosX < -float64(cloud.ImgW)-2000 {
		cloud.init()
		cloud.PosX = config.WindowWidth / cloud.Scale
	}
	if forecast.WindDeg >= 0 && forecast.WindDeg < 180 {
		cloud.PosX = cloud.PosX + cloud.VelX*config.WindSpeedModifier
	}
	if forecast.WindDeg >= 180 && forecast.WindDeg < 360 {
		cloud.PosX = cloud.PosX - cloud.VelX*config.WindSpeedModifier
	}
}

func (cloud *Cloud) Draw(screen *ebiten.Image) {
	cloud.op.GeoM.Reset()
	cloud.op.GeoM.Translate(
		cloud.PosX,
		cloud.PosY,
	)
	cloud.op.GeoM.Scale(
		cloud.Scale,
		cloud.Scale,
	)
	cloud.op.Filter = ebiten.FilterLinear
	screen.DrawImage(cloud.Texture, &cloud.op)
}

type Clouds struct {
	Clouds []*Cloud
	Num    int
}

func (clouds *Clouds) Update(forecast models.DailyForecast) {
	for i := 0; i < clouds.Num; i++ {
		clouds.Clouds[i].Update(forecast)
	}
}

func (clouds *Clouds) Draw(screen *ebiten.Image) {
	for i := 0; i < clouds.Num; i++ {
		clouds.Clouds[i].Draw(screen)
	}
}
