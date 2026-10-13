package engine

type Entity struct {
	X, Y   float64
	Sprite *Sprite
	Speed  float64
}

func (e *Entity) Move(dx, dy float64) {
	e.X += dx
	e.Y += dy
}
