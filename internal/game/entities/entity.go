package entities

import "github.com/hajimehoshi/ebiten/v2"

type Entity struct {
	X, Y   int
	Sprite *ebiten.Image
}

func (e *Entity) move(dx, dy int) {
	e.X += dx
	e.Y += dy
}
