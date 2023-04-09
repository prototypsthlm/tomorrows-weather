package main

import (
	"embed"
	"fmt"
	"image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sandnuggah/tomorrows-weather/config"
	"github.com/sandnuggah/tomorrows-weather/game"
)

//go:embed shaders/snow.kage
var snowShader []byte

//go:embed shaders/stars.kage
var starsShader []byte

//go:embed textures/*
var textures embed.FS

var (
	shaders        []*ebiten.Shader
	aCloudTextures []*ebiten.Image
	bCloudTextures []*ebiten.Image
	cCloudTextures []*ebiten.Image
	dCloudTextures []*ebiten.Image
	skyTexture     *ebiten.Image
)

func init() {
	snowShader, err := ebiten.NewShader(snowShader)
	if err != nil {
		log.Fatal(err)
	}
	starsShader, err := ebiten.NewShader(starsShader)
	if err != nil {
		log.Fatal(err)
	}
	shaders = append(shaders, snowShader, starsShader)
}

func init() {
	raw, err := textures.Open("textures/sky.png")
	if err != nil {
		log.Fatal(err)
	}
	decoded, err := png.Decode(raw)
	if err != nil {
		log.Fatal(err)
	}
	skyTexture = ebiten.NewImageFromImage(decoded)
}

func init() {
	for i := 1; i <= 1; i++ {
		raw, err := textures.Open(fmt.Sprintf("textures/1/%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		decoded, err := png.Decode(raw)
		if err != nil {
			log.Fatal(err)
		}
		image := ebiten.NewImageFromImage(decoded)
		texture := ebiten.NewImage(image.Size())
		texture.DrawImage(
			image,
			&ebiten.DrawImageOptions{},
		)
		aCloudTextures = append(aCloudTextures, texture)
	}
}

func init() {
	for i := 1; i <= 5; i++ {
		raw, err := textures.Open(fmt.Sprintf("textures/2/%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		decoded, err := png.Decode(raw)
		if err != nil {
			log.Fatal(err)
		}
		image := ebiten.NewImageFromImage(decoded)
		texture := ebiten.NewImage(image.Size())
		texture.DrawImage(
			image,
			&ebiten.DrawImageOptions{},
		)
		bCloudTextures = append(bCloudTextures, texture)
	}
}
func init() {
	for i := 1; i <= 4; i++ {
		raw, err := textures.Open(fmt.Sprintf("textures/3/%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		decoded, err := png.Decode(raw)
		if err != nil {
			log.Fatal(err)
		}
		image := ebiten.NewImageFromImage(decoded)
		texture := ebiten.NewImage(image.Size())
		texture.DrawImage(
			image,
			&ebiten.DrawImageOptions{},
		)
		cCloudTextures = append(cCloudTextures, texture)
	}
}
func init() {
	for i := 1; i <= 3; i++ {
		raw, err := textures.Open(fmt.Sprintf("textures/4/%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		decoded, err := png.Decode(raw)
		if err != nil {
			log.Fatal(err)
		}
		image := ebiten.NewImageFromImage(decoded)
		texture := ebiten.NewImage(image.Size())
		texture.DrawImage(
			image,
			&ebiten.DrawImageOptions{},
		)
		dCloudTextures = append(dCloudTextures, texture)
	}
}

func main() {
	ebiten.SetWindowSize(
		config.WindowWidth,
		config.WindowHeight,
	)
	ebiten.SetWindowTitle("Tomorrows Weather")
	if err := ebiten.RunGame(&game.Game{
		Shaders:        shaders,
		ACloudTextures: aCloudTextures,
		BCloudTextures: bCloudTextures,
		CCloudTextures: cCloudTextures,
		DCloudTextures: dCloudTextures,
		SkyTexture:     skyTexture,
	}); err != nil {
		log.Fatal(err)
	}
}
