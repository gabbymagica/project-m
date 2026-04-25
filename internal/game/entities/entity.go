package entities

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	X, Y   float64
	Sprite *ebiten.Image
}
