package models

import "github.com/faiface/pixel"

type Cloud struct {
	Sprite         *pixel.Sprite //sprite to render
	PositionVec    pixel.Vec     //current position to render
	AnimationDelta float64
	PositionY      float64 //rand float for dynamic animations
	PositionX      float64
}
