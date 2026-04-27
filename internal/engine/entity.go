package engine

type Entity struct {
	X, Y   int
	Sprite *Sprite
}

func (e *Entity) move(dx, dy int) {
	e.X = dx
	e.Y = dy
	e.Sprite.X = dx
	e.Sprite.Y = dy
}
