package game

import (
	_ "embed"
	_ "image/png"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sandnuggah/tomorrows-weather/api"
	"github.com/sandnuggah/tomorrows-weather/config"
	"github.com/sandnuggah/tomorrows-weather/models"
	"github.com/sandnuggah/tomorrows-weather/sprites"
	"github.com/sandnuggah/tomorrows-weather/utils"
)

// Timișoara, Romania
var (
	lat = 45.7489
	lon = 21.2087
)

type Game struct {
	Shaders        []*ebiten.Shader
	SkyTexture     *ebiten.Image
	FogTexture     *ebiten.Image
	ACloudTextures []*ebiten.Image
	BCloudTextures []*ebiten.Image
	CCloudTextures []*ebiten.Image
	DCloudTextures []*ebiten.Image
	ECloudTextures []*ebiten.Image

	inited            bool
	snowAmount        int
	timezone          string
	location          *time.Location
	forecast          models.DailyForecast
	time              int
	lastWeatherUpdate time.Time
	skyImage          *ebiten.Image
	fogImage          *ebiten.Image
	skySaturation     float64
	skyBrightness     float64
	cloudOpacity      float64
	isFoggy           bool
	sprites           struct {
		aClouds   sprites.Clouds
		bClouds   sprites.Clouds
		cClouds   sprites.Clouds
		dClouds   sprites.Clouds
		eClouds   sprites.Clouds
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

	// ------------------------------------------------
	weatherId := config.DefaultWeatherId
	if len(game.forecast.Weather) > 0 {
		weatherId = game.forecast.Weather[0].ID
	}
	game.sprites.aClouds.Num,
		game.sprites.bClouds.Num,
		game.sprites.cClouds.Num,
		game.sprites.dClouds.Num,
		game.sprites.eClouds.Num,
		game.sprites.raindrops.Num,
		game.snowAmount,
		game.skySaturation,
		game.skyBrightness,
		game.cloudOpacity,
		game.isFoggy =
		utils.WeatherConditionIdToConfig(weatherId)

	game.skyImage = utils.DrawSky(game.SkyTexture, game.location)
	game.fogImage = utils.DrawFog(game.FogTexture)
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
			ImgW:    w,
			ImgH:    h,
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.bClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.bClouds.Clouds {
		texture := game.BCloudTextures[utils.Rand(0, len(game.BCloudTextures))]
		w, h := texture.Size()
		game.sprites.bClouds.Clouds[i] = &sprites.Cloud{
			ImgW:    w,
			ImgH:    h,
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.cClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.cClouds.Clouds {
		texture := game.CCloudTextures[utils.Rand(0, len(game.CCloudTextures))]
		w, h := texture.Size()
		game.sprites.cClouds.Clouds[i] = &sprites.Cloud{
			ImgW:    w,
			ImgH:    h,
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.dClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.dClouds.Clouds {
		texture := game.DCloudTextures[utils.Rand(0, len(game.DCloudTextures))]
		w, h := texture.Size()
		game.sprites.dClouds.Clouds[i] = &sprites.Cloud{
			ImgW:    w,
			ImgH:    h,
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.eClouds.Clouds = make(
		[]*sprites.Cloud,
		config.MaxClouds,
	)
	for i := range game.sprites.eClouds.Clouds {
		texture := game.ECloudTextures[utils.Rand(0, len(game.ECloudTextures))]
		w, h := texture.Size()
		game.sprites.eClouds.Clouds[i] = &sprites.Cloud{
			ImgW:    w,
			ImgH:    h,
			Texture: texture,
		}
	}
	// ------------------------------------------------
	game.sprites.raindrops.Raindrops = make(
		[]*sprites.Raindrop,
		config.MaxRaindrops,
	)
	for i := range game.sprites.raindrops.Raindrops {
		game.sprites.raindrops.Raindrops[i] = &sprites.Raindrop{}
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
			game.sprites.eClouds.Num,
			game.sprites.raindrops.Num,
			game.snowAmount,
			game.skySaturation,
			game.skyBrightness,
			game.cloudOpacity,
			game.isFoggy =
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
	game.sprites.cClouds.Num = utils.Clamp(
		game.sprites.cClouds.Num,
		0,
		config.MaxClouds,
	)
	game.sprites.dClouds.Num = utils.Clamp(
		game.sprites.dClouds.Num,
		0,
		config.MaxClouds,
	)
	game.sprites.eClouds.Num = utils.Clamp(
		game.sprites.eClouds.Num,
		0,
		config.MaxClouds,
	)
	game.sprites.raindrops.Num = utils.Clamp(
		game.sprites.raindrops.Num,
		0,
		config.MaxRaindrops,
	)

	// Set the sky saturation and cloud opacity based on the time of day
	co := game.cloudOpacity

	switch time.Now().In(game.location).Hour() {
	case 20, 21, 22, 23, 24, 0, 1, 2, 3, 4:
		co = game.cloudOpacity / 6
		game.skySaturation = 0.5
	}

	game.sprites.aClouds.Update(game.forecast, co)
	game.sprites.bClouds.Update(game.forecast, co)
	game.sprites.cClouds.Update(game.forecast, co)
	game.sprites.dClouds.Update(game.forecast, co)
	game.sprites.eClouds.Update(game.forecast, co)

	game.sprites.raindrops.Update(game.forecast)

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	starsIntensity := 0.0
	skyOp := &ebiten.DrawImageOptions{}
	fogOp := &ebiten.DrawImageOptions{}

	skyOp.ColorM.Reset()
	fogOp.ColorM.Reset()

	switch time.Now().In(game.location).Hour() {
	case 21:
		skyOp.ColorM.ChangeHSV(0, 0.9, game.skyBrightness)
		starsIntensity = 0.1
		fogOp.CompositeMode = ebiten.CompositeModeXor
	case 22:
		skyOp.ColorM.ChangeHSV(0, 0.6, game.skyBrightness)
		starsIntensity = 0.3
		fogOp.CompositeMode = ebiten.CompositeModeXor
	case 23:
		skyOp.ColorM.ChangeHSV(0, 0.4, game.skyBrightness)
		starsIntensity = 0.6
		fogOp.CompositeMode = ebiten.CompositeModeXor
	case 0, 1, 2, 3:
		skyOp.ColorM.ChangeHSV(0, 0.2, game.skyBrightness)
		starsIntensity = 1.0
		fogOp.CompositeMode = ebiten.CompositeModeXor
	case 4:
		skyOp.ColorM.ChangeHSV(0, 0.4, game.skyBrightness)
		starsIntensity = 0.6
		fogOp.CompositeMode = ebiten.CompositeModeXor
	case 5:
		skyOp.ColorM.ChangeHSV(0, 0.6, game.skyBrightness)
		starsIntensity = 0.3
		fogOp.CompositeMode = ebiten.CompositeModeXor
	case 6:
		skyOp.ColorM.ChangeHSV(0, 0.9, game.skyBrightness)
		starsIntensity = 0.1
		fogOp.CompositeMode = ebiten.CompositeModeXor
	default:
		skyOp.ColorM.ChangeHSV(0, 1, game.skyBrightness)
		fogOp.ColorM.ChangeHSV(1, 1, 0.5)
		fogOp.CompositeMode = ebiten.CompositeModeSourceAtop
		starsIntensity = 0.0
	}

	screen.DrawImage(
		game.skyImage,
		skyOp,
	)

	if game.isFoggy {
		screen.DrawImage(
			game.fogImage,
			fogOp,
		)
	}

	screen.DrawRectShader(
		config.WindowWidth,
		config.WindowHeight,
		game.Shaders[1], // stars
		&ebiten.DrawRectShaderOptions{
			Uniforms: map[string]interface{}{
				"Time":      float32(game.time) / 60,
				"Intensity": float32(starsIntensity),
			},
		},
	)

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
	game.sprites.cClouds.Draw(screen)
	game.sprites.dClouds.Draw(screen)
	game.sprites.eClouds.Draw(screen)

	game.sprites.raindrops.Draw(screen)

	// ebitenutil.DebugPrint(
	// 	screen,
	// 	fmt.Sprintf(
	// 		"code=%d tps=%f fps=%f time=%s",
	// 		game.forecast.Weather[0].ID,
	// 		ebiten.ActualTPS(),
	// 		ebiten.ActualFPS(),
	// 		time.Now().In(game.location).Format("15:04:05"),
	// 	),
	// )
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return config.WindowWidth, config.WindowHeight
}
