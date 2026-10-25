package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	ID     string
	Name   string
	Image  *ebiten.Image
	Alpha  float32
	drawOp ebiten.DrawImageOptions
}

func NewSprite(id, name string, image *ebiten.Image) *Sprite {
	return &Sprite{
		ID:     id,
		Name:   name,
		Image:  image,
		Alpha:  1.0,
		drawOp: ebiten.DrawImageOptions{},
	}
}

func (s *Sprite) Instantiate() *Sprite {
	newSprite := *s

	return &newSprite
}

func (s *Sprite) Draw(screen *ebiten.Image, x, y float64, scaleX, scaleY float64, alpha float32) {
	if s == nil || s.Image == nil || screen == nil {
		return
	}

	s.drawOp.GeoM.Reset()
	s.drawOp.ColorScale.Reset()

	s.drawOp.ColorScale.ScaleAlpha(alpha)
	s.drawOp.GeoM.Scale(scaleX, scaleY)

	s.drawOp.GeoM.Translate(x, y)
	screen.DrawImage(s.Image, &s.drawOp)
}

type SpriteManager struct {
	Sprites map[string]*Sprite
}

func NewSpriteManager() *SpriteManager {
	return &SpriteManager{
		Sprites: make(map[string]*Sprite),
	}
}

func (sm *SpriteManager) AddSprite(sprite *Sprite) {
	sm.Sprites[sprite.ID] = sprite
}
