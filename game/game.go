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
	lat = 34.0522
	lon = 118.2437
	// lat = 45.7489
	// lon = 21.2087
)

type Game struct {
	Shaders        []*ebiten.Shader
	SkyTexture     *ebiten.Image
	ACloudTextures []*ebiten.Image
	BCloudTextures []*ebiten.Image
	CCloudTextures []*ebiten.Image
	DCloudTextures []*ebiten.Image

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
		aClouds   sprites.Clouds
		bClouds   sprites.Clouds
		cClouds   sprites.Clouds
		dClouds   sprites.Clouds
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
	// ------------------------------------------------
	game.sprites.aClouds.Num,
		game.sprites.bClouds.Num,
		game.sprites.cClouds.Num,
		game.sprites.dClouds.Num,
		game.sprites.raindrops.Num,
		game.snowAmount,
		game.skySaturation,
		game.skyBrightness =
		utils.WeatherConditionIdToConfig(weatherId)

	game.skyImage = utils.DrawSky(game.SkyTexture, game.location)
	game.lastWeatherUpdate = time.Now()

	// ------------------------------------------------
	game.sprites.aClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.aClouds.Clouds {
		texture := game.ACloudTextures[0]
		w, h := texture.Size()
		game.sprites.aClouds.Clouds[i] = &sprites.Cloud{
			ImgW: w,
			ImgH: h,
			PosX: utils.Randf(-1500, config.WindowWidth+1500),
			PosY: utils.Randf(
				200,
				config.WindowHeight+500,
			),
			VelX:    utils.Randf(0.5, 1.5),
			Scale:   utils.Randf(0.6, 0.9),
			Opacity: utils.Randf(0.5, 0.8),
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.bClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.bClouds.Clouds {
		texture := game.BCloudTextures[utils.Rand(1, 5)]
		w, h := texture.Size()
		game.sprites.bClouds.Clouds[i] = &sprites.Cloud{
			ImgW: w,
			ImgH: h,
			PosX: utils.Randf(-1500, config.WindowWidth+1500),
			PosY: utils.Randf(
				200,
				config.WindowHeight+500,
			),
			VelX:    utils.Randf(0.5, 1.5),
			Scale:   utils.Randf(0.6, 0.9),
			Opacity: utils.Randf(0.5, 0.8),
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.cClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.cClouds.Clouds {
		texture := game.BCloudTextures[utils.Rand(1, 4)]
		w, h := texture.Size()
		game.sprites.cClouds.Clouds[i] = &sprites.Cloud{
			ImgW: w,
			ImgH: h,
			PosX: utils.Randf(-1500, config.WindowWidth+1500),
			PosY: utils.Randf(
				200,
				config.WindowHeight+500,
			),
			VelX:    utils.Randf(0.5, 1.5),
			Scale:   utils.Randf(0.6, 0.9),
			Opacity: utils.Randf(0.5, 0.8),
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.dClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.dClouds.Clouds {
		texture := game.DCloudTextures[utils.Rand(1, 3)]
		w, h := texture.Size()
		game.sprites.dClouds.Clouds[i] = &sprites.Cloud{
			ImgW: w,
			ImgH: h,
			PosX: utils.Randf(-1500, config.WindowWidth+1500),
			PosY: utils.Randf(
				200,
				config.WindowHeight+500,
			),
			VelX:    utils.Randf(0.5, 1.5),
			Scale:   utils.Randf(0.6, 0.9),
			Opacity: utils.Randf(0.5, 0.8),
			Texture: texture,
		}
	}
	// ------------------------------------------------
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

		game.sprites.aClouds.Num,
			game.sprites.bClouds.Num,
			game.sprites.cClouds.Num,
			game.sprites.dClouds.Num,
			game.sprites.raindrops.Num,
			game.snowAmount,
			game.skySaturation,
			game.skyBrightness =
			utils.WeatherConditionIdToConfig(weatherId)
		game.skyImage = utils.DrawSky(game.SkyTexture, game.location)
		game.lastWeatherUpdate = time.Now()
	}
	game.sprites.aClouds.Num = utils.Clamp(
		game.sprites.aClouds.Num,
		0,
		config.MaxClouds,
	)
	game.sprites.bClouds.Num = utils.Clamp(
		game.sprites.bClouds.Num,
		0,
		config.MaxClouds,
	)
	game.sprites.raindrops.Num = utils.Clamp(
		game.sprites.raindrops.Num,
		0,
		config.MaxRaindrops,
	)
	game.sprites.aClouds.Update(game.forecast)
	game.sprites.bClouds.Update(game.forecast)
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
	default:
		screen.DrawRectShader(
			config.WindowWidth,
			config.WindowHeight,
			game.Shaders[1], // stars
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
	game.sprites.aClouds.Draw(screen)
	game.sprites.bClouds.Draw(screen)
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
