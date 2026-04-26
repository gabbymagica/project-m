package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	ID             string
	Name           string
	Image          *ebiten.Image
	ScaleX, ScaleY float64
	X, Y           int
}

func NewSprite(id, name string, image *ebiten.Image) *Sprite {
	return &Sprite{
		ID:     id,
		Name:   name,
		Image:  image,
		ScaleX: 1.0,
		ScaleY: 1.0,
	}
}

func (s *Sprite) Instantiate() *Sprite {
	newSprite := *s

	return &newSprite
}

func (s *Sprite) Scale(factorX, factorY float64) {
	s.ScaleX = factorX
	s.ScaleY = factorY
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	if s == nil || s.Image == nil || screen == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(s.ScaleX, s.ScaleY)

	op.GeoM.Translate(float64(s.X), float64(s.Y))
	screen.DrawImage(s.Image, op)
}
