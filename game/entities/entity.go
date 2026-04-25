package entities

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	X, Y   float64
	Sprite *ebiten.Image
}

func (e *Entity) move(dx, dy float64) {
	e.X += dx
	e.Y += dy
}