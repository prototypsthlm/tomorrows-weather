package main

import (
	_ "embed"
	"fmt"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/sandnuggah/tomorrows-weather/config"
	"github.com/sandnuggah/tomorrows-weather/game"
)

//go:embed shaders/snow.kage
var snowShader []byte

//go:embed shaders/gradient.kage
var gradientShader []byte

//go:embed shaders/stars.kage
var starsShader []byte

var (
	shaders       []*ebiten.Shader
	cloudTextures []*ebiten.Image
	skyTexture    *ebiten.Image
)

func init() {
	snowShader, err := ebiten.NewShader(snowShader)
	if err != nil {
		log.Fatal(err)
	}
	gradientShader, err := ebiten.NewShader(gradientShader)
	if err != nil {
		log.Fatal(err)
	}
	starsShader, err := ebiten.NewShader(starsShader)
	if err != nil {
		log.Fatal(err)
	}
	shaders = append(shaders, snowShader, gradientShader, starsShader)
}

func init() {
	image, _, err := ebitenutil.NewImageFromFile("textures/sky.png")
	if err != nil {
		log.Fatal(err)
	}
	skyTexture = image
}

func init() {
	for i := 1; i <= 5; i++ {
		image, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("textures/%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		texture := ebiten.NewImage(image.Size())
		texture.DrawImage(
			image,
			&ebiten.DrawImageOptions{},
		)
		cloudTextures = append(cloudTextures, texture)
	}
}

func main() {
	ebiten.SetWindowSize(
		config.WindowWidth,
		config.WindowHeight,
	)
	ebiten.SetWindowTitle("Tomorrows Weather")
	if err := ebiten.RunGame(&game.Game{
		Shaders:       shaders,
		CloudTextures: cloudTextures,
		SkyTexture:    skyTexture,
	}); err != nil {
		log.Fatal(err)
	}
}
