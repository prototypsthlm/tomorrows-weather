package game

import (
	_ "embed"
	"fmt"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sandnuggah/tomorrows-weather/api"
	"github.com/sandnuggah/tomorrows-weather/config"
	"github.com/sandnuggah/tomorrows-weather/models"
	"github.com/sandnuggah/tomorrows-weather/sprites"
	"github.com/sandnuggah/tomorrows-weather/utils"
)

var (
	lat = 45.7489
	lon = 21.2087
)

type Game struct {
	Shaders       []*ebiten.Shader
	SkyTexture    *ebiten.Image
	CloudTextures []*ebiten.Image

	inited            bool
	snowAmount        int
	timezone          string
	location          *time.Location
	forecast          models.DailyForecast
	time              int
	lastWeatherUpdate time.Time
	skyImage          *ebiten.Image
	skySaturation     float64
	skyBrightness     float64
	sprites           struct {
		clouds    sprites.Clouds
		raindrops sprites.Raindrops
	}
}

func (game *Game) init() {
	defer func() {
		game.inited = true
	}()
	forecast, timezone := api.GetForecast(lon, lat)
	game.forecast = forecast
	game.timezone = timezone
	game.location, _ = time.LoadLocation(game.timezone)

	weatherId := config.DefaultWeatherId
	if len(game.forecast.Weather) > 0 {
		weatherId = game.forecast.Weather[0].ID
	}

	game.sprites.clouds.Num,
		game.sprites.raindrops.Num,
		game.snowAmount,
		game.skySaturation,
		game.skyBrightness =
		utils.WeatherConditionIdToConfig(weatherId)

	game.skyImage = utils.DrawSky(game.SkyTexture, game.location)
	game.lastWeatherUpdate = time.Now()
	game.sprites.clouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.clouds.Clouds {
		texture := game.CloudTextures[utils.Rand(1, 5)]
		w, h := texture.Size()
		game.sprites.clouds.Clouds[i] = &sprites.Cloud{
			ImgW: w,
			ImgH: h,
			PosX: utils.Randf(-1500, config.WindowWidth+1500),
			PosY: utils.Randf(
				-200,
				config.WindowHeight+200,
			),
			VelX:    utils.Randf(0.5, 1.5),
			Scale:   utils.Randf(0.6, 0.9),
			Opacity: utils.Randf(0.5, 0.8),
			Texture: texture,
		}
	}
	game.sprites.raindrops.Raindrops = make(
		[]*sprites.Raindrop,
		config.MaxRaindrops,
	)
	for i := range game.sprites.raindrops.Raindrops {
		game.sprites.raindrops.Raindrops[i] = &sprites.Raindrop{
			VelX:  utils.Randf(-0.5, 0.5),
			VelY:  utils.Randf(15, 20),
			Alpha: utils.Rand(50, 127),
		}
	}
}

func (game *Game) Update() error {
	if !game.inited {
		game.init()
	}
	game.time++
	if game.lastWeatherUpdate.Before(
		time.Now().Add(config.UpdateWeatherInterval),
	) {
		forecast, timezone := api.GetForecast(lon, lat)
		game.forecast = forecast
		game.timezone = timezone
		game.location, _ = time.LoadLocation(game.timezone)

		weatherId := config.DefaultWeatherId
		if len(game.forecast.Weather) > 0 {
			weatherId = game.forecast.Weather[0].ID
		}

		game.sprites.clouds.Num,
			game.sprites.raindrops.Num,
			game.snowAmount,
			game.skySaturation,
			game.skyBrightness =
			utils.WeatherConditionIdToConfig(weatherId)
		game.skyImage = utils.DrawSky(game.SkyTexture, game.location)
		game.lastWeatherUpdate = time.Now()
	}
	game.sprites.clouds.Num = utils.Clamp(
		game.sprites.clouds.Num,
		0,
		config.MaxClouds,
	)
	game.sprites.raindrops.Num = utils.Clamp(
		game.sprites.raindrops.Num,
		0,
		config.MaxRaindrops,
	)
	game.sprites.clouds.Update(game.forecast)
	game.sprites.raindrops.Update(game.forecast)
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorM.ChangeHSV(0, game.skySaturation, 1)
	screen.DrawImage(
		game.skyImage,
		op,
	)
	switch time.Now().In(game.location).Hour() {
	case 22, 23, 0, 1, 2, 3, 4:
		screen.DrawRectShader(
			config.WindowWidth,
			config.WindowHeight,
			game.Shaders[2], // stars
			&ebiten.DrawRectShaderOptions{
				Uniforms: map[string]interface{}{
					"Time": float32(game.time) / 60,
				},
			},
		)
	}
	screen.DrawRectShader(
		config.WindowWidth,
		config.WindowHeight,
		game.Shaders[0], // snow
		&ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"Time":   float32(game.time) / 60,
				"Depth":  float32(0.7),
				"Width":  float32(0.3),
				"Speed":  float32(2.6),
				"Amount": float32(game.snowAmount),
			},
			CompositeMode: ebiten.CompositeModeLighter,
		},
	)
	game.sprites.clouds.Draw(screen)
	game.sprites.raindrops.Draw(screen)
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf(
			"code=%d tps=%f fps=%f",
			500,
			ebiten.ActualTPS(),
			ebiten.ActualFPS(),
		),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.WindowWidth, config.WindowHeight
}
