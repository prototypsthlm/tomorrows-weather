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
	cloudOpacity = 1
	isFoggy = false

	switch id {
	// thunderstorm with light rain
	case 200:
		raindropsNum = 50
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// thunderstorm with rain
	case 201:
		raindropsNum = 100
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// thunderstorm with heavy rain
	case 202:
		raindropsNum = 200
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// 210: light thunderstorm
	case 210:
		raindropsNum = 50
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// 211: thunderstorm
	case 211:
		raindropsNum = 50
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// 212: heavy thunderstorm
	case 212:
		raindropsNum = 50
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// 221: ragged thunderstorm
	case 221:
		raindropsNum = 50
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// 230: thunderstorm with light drizzle
	case 230:
		raindropsNum = 50
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// thunderstorm with drizzle
	case 231:
		raindropsNum = 50
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// thunderstorm with heavy drizzle
	case 232:
		raindropsNum = 100
		skySaturation = 0.2
		cloudOpacity = 0.5
		isFoggy = true

	// light intensity drizzle
	case 300:
		raindropsNum = 50
		skySaturation = 0.2
		isFoggy = true

	// drizzle
	case 301:
		raindropsNum = 50
		skySaturation = 0.2
		isFoggy = true

	// heavy intensity drizzle
	case 302:
		raindropsNum = 100
		skySaturation = 0.2
		isFoggy = true

	// light intensity drizzle rain
	case 310:
		raindropsNum = 100
		skySaturation = 0.2
		isFoggy = true

	// drizzle rain
	case 311:
		raindropsNum = 100
		skySaturation = 0.2
		isFoggy = true

	// heavy intensity drizzle rain
	case 312:
		raindropsNum = 150
		skySaturation = 0.2
		isFoggy = true

	// shower rain and drizzle
	case 313:
		raindropsNum = 250
		skySaturation = 0.2
		isFoggy = true

	// heavy shower rain and drizzle
	case 314:
		raindropsNum = 250
		skySaturation = 0.2
		isFoggy = true

	// shower drizzle
	case 321:
		raindropsNum = 250
		skySaturation = 0.2
		isFoggy = true

	// light rain
	case 500:
		raindropsNum = 50
		skySaturation = 0.2
		isFoggy = true

	// moderate rain
	case 501:
		raindropsNum = 100
		skySaturation = 0.2
		isFoggy = true

	// heavy intensity rain
	case 502:
		raindropsNum = 250
		skySaturation = 0.2
		isFoggy = true

	// very heavy rain
	case 503:
		raindropsNum = 450
		skySaturation = 0.2
		isFoggy = true

	// extreme rain
	case 504:
		raindropsNum = 450
		skySaturation = 0.2
		isFoggy = true

	// freezing rain
	case 511:
		raindropsNum = 450
		skySaturation = 0.2
		isFoggy = true

	// light intensity shower rain
	case 520:
		raindropsNum = 50
		skySaturation = 0.2
		isFoggy = true

	// shower rain
	case 521:
		raindropsNum = 200
		skySaturation = 0.2
		isFoggy = true

	// heavy intensity shower rain
	case 522:
		raindropsNum = 450
		skySaturation = 0.2
		isFoggy = true

	// ragged shower rain
	case 531:
		raindropsNum = 450
		skySaturation = 0.2
		isFoggy = true

	// light snow
	case 600:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 10

	// snow
	case 601:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 20

	// heavy snow
	case 602:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 20

	// sleet
	case 611:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 5
		raindropsNum = 50

	// light shower sleet
	case 612:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 5
		raindropsNum = 50

	// shower sleet
	case 613:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 5
		raindropsNum = 50

	// light rain and snow
	case 615:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 10
		raindropsNum = 50

	// rain and snow
	case 616:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 10
		raindropsNum = 50

	// light shower snow
	case 620:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 20
		raindropsNum = 100

	// shower snow
	case 621:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 20
		raindropsNum = 100

	// heavy shower snow
	case 622:
		skySaturation = 0.2
		isFoggy = true
		snowAmount = 30
		raindropsNum = 100

	// clear
	case 800:

	// few clouds 11-25%
	case 801:
		aCloudsNum = 1
		cloudOpacity = 0.5

	// scattered clouds 26-50%
	case 802:
		aCloudsNum = 2
		cloudOpacity = 0.5

	// broken clouds 51-84%
	case 803:
		aCloudsNum = 2
		bCloudsNum = 3
		cloudOpacity = 0.5

	// overcast clouds 85-100%
	case 804:
		skySaturation = 0.5
		cloudOpacity = 0.2
		isFoggy = true
		cCloudsNum = 3
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
