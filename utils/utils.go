package utils

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sandnuggah/tomorrows-weather/config"
)

// Rand returns a random int between [min, max]s
func Rand(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// Randf returns a random float between [min, max]
func Randf(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Float64()*(max-min)
}

// Clamp returns f clamped to [low, high]
func Clamp(i, low, high int) int {
	if i < low {
		return low
	}
	if i > high {
		return high
	}
	return i
}

// InTimeSpan returns true if [check] is within [start, end]
func InTimeSpan(start, end, check time.Time) bool {
	_end := end
	_check := check
	if end.Before(start) {
		_end = end.Add(24 * time.Hour)
		if check.Before(start) {
			_check = check.Add(24 * time.Hour)
		}
	}
	return _check.After(start) && _check.Before(_end)
}

// Scale returns [n] scaled to [min, max]
func Scale(min, max, n float64) float64 {
	return (n - min) / (max - min)
}

// Bod returns beginning of day
func Bod(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

// DrawSky returns a vertical slice of skyTexture covering the entire screen
func DrawSky(skyTexture *ebiten.Image, location *time.Location) *ebiten.Image {
	img := ebiten.NewImage(
		config.WindowWidth,
		config.WindowHeight,
	)
	hour := time.Now().In(location).Hour()

	for i := 0; i < config.WindowHeight; i++ {
		pix := skyTexture.At(
			int(math.Abs(12.5*float64(hour))),
			i,
		)
		for j := 0; j < config.WindowWidth; j++ {
			img.Set(j, i, pix)
		}
	}
	return img
}

func DrawFog(fogTexture *ebiten.Image) *ebiten.Image {
	return ebiten.NewImageFromImage(fogTexture)
}

// WeatherConditionIdToConfig returns number of raindrops and number of clouds
func WeatherConditionIdToConfig(id int) (aCloudsNum, bCloudsNum, cCloudsNum, dCloudsNum, raindropsNum, snowAmount int, skySaturation, skyBrightness, cloudOpacity float64, isFoggy bool) {
	aCloudsNum = 0
	bCloudsNum = 0
	cCloudsNum = 0
	dCloudsNum = 0
	snowAmount = 0
	raindropsNum = 0
	skySaturation = 1
	skyBrightness = 1
	cloudOpacity = 0.5
	isFoggy = true

	id = 802

	switch id {
	// thunderstorm with light rain
	case 200:
		aCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 50, 0.2, 1

	// thunderstorm with rain
	case 201:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 100, 0.2, 1

	// thunderstorm with heavy rain
	case 202:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 250, 0.2, 1

	// 210: light thunderstorm
	// 211: thunderstorm
	// 212: heavy thunderstorm
	// 221: ragged thunderstorm
	// 230: thunderstorm with light drizzle
	case 210:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 50, 0.2, 1
	case 211:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 50, 0.2, 1
	case 212:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 50, 0.2, 1
	case 221:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 50, 0.2, 1
	case 230:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 50, 0.2, 1

	// thunderstorm with drizzle
	case 231:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 30, 0.2, 1

	// thunderstorm with heavy drizzle
	case 232:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 60, 0.2, 1

	// light intensity drizzle
	case 300:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 100, 0.2, 1

	// drizzle
	case 301:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 100, 0.2, 1

	// heavy intensity drizzle
	case 302:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 150, 0.2, 1

	// light intensity drizzle rain
	case 310:
		dCloudsNum, raindropsNum, skySaturation, skyBrightness = 50, 50, 0.2, 1

	// drizzle rain
	case 311:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 100, 2, 0.9, 1

	// heavy intensity drizzle rain
	case 312:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 350, 2, 0.9, 1

	// shower rain and drizzle
	case 313:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 350, 2, 0.9, 1

	// heavy shower rain and drizzle
	case 314:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 512, 2, 0.9, 1

	// shower drizzle
	case 321:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 70, 2, 0.9, 1

	// light rain
	case 500:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 70, 4, 0.3, 1

	// moderate rain
	case 501:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 100, 4, 0.9, 1

	// heavy intensity rain
	case 502:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 512, 4, 0.9, 1

	// very heavy rain
	case 503:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 512, 4, 0.9, 1

	// extreme rain
	case 504:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 512, 4, 0.5, 1

	// freezing rain
	case 511:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 256, 4, 0.5, 1

	// light intensity shower rain
	case 520:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 256, 4, 0.8, 1

	// shower rain
	case 521:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 256, 4, 0.6, 1

	// heavy intensity shower rain
	case 522:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 468, 4, 0.5, 1

	// ragged shower rain
	case 531:
		raindropsNum, aCloudsNum, skySaturation, skyBrightness = 512, 4, 0.5, 1

	// light snow
	case 600:
		snowAmount, aCloudsNum, skySaturation, skyBrightness = 10, 4, 0.5, 1

	// snow
	case 601:
		snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 4, 0.5, 1

	// heavy snow
	case 602:
		snowAmount, aCloudsNum, skySaturation, skyBrightness = 30, 4, 0.5, 1

	// sleet
	case 611:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 10, 0, 0.5, 1

	// light shower sleet
	case 612:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 10, 0, 0.5, 1

	// shower sleet
	case 613:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 10, 0, 0.5, 1

	// light rain and snow
	case 615:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 10, 0, 0.5, 1

	// rain and snow
	case 616:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 10, 0, 0.5, 1

	// light shower snow
	case 620:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 10, 0, 0.5, 1

	// shower snow
	case 621:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 20, 0, 0.5, 1

	// heavy shower snow
	case 622:
		raindropsNum, snowAmount, aCloudsNum, skySaturation, skyBrightness = 20, 30, 0, 0.5, 1

	// clear
	case 800:
		aCloudsNum, skySaturation, skyBrightness = 0, 0.8, 1

	// few clouds 11-25%
	case 801:
		aCloudsNum, skySaturation, skyBrightness = 2, 1, 1

	// scattered clouds 26-50%
	case 802:
		aCloudsNum = 0
		bCloudsNum = 4
		cCloudsNum = 0
		dCloudsNum = 0
		skySaturation, skyBrightness = 1, 1

	// broken clouds 51-84%
	case 803:
		bCloudsNum, bCloudsNum, skySaturation, skyBrightness = 4, 12, 0.5, 0.7

	// overcast clouds 85-100%
	case 804:
		aCloudsNum, skySaturation, skyBrightness = 12, 0.5, 0.7
	}

	return aCloudsNum,
		bCloudsNum,
		cCloudsNum,
		dCloudsNum,
		raindropsNum,
		snowAmount,
		skySaturation,
		skyBrightness,
		cloudOpacity,
		isFoggy
}
