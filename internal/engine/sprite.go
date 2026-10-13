package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	ID             string
	Name           string
	Image          *ebiten.Image
	Alpha          float32
}

func NewSprite(id, name string, image *ebiten.Image) *Sprite {
	return &Sprite{
		ID:     id,
		Name:   name,
		Image:  image,
		Alpha:  1.0,
	}
}

func (s *Sprite) Instantiate() *Sprite {
	newSprite := *s

	return &newSprite
}

func (s *Sprite) SetAlpha(alpha float32) {
	s.Alpha = alpha
}

func (s *Sprite) Draw(screen *ebiten.Image, x, y float64, scaleX, scaleY float64, ) {
	if s == nil || s.Image == nil || screen == nil {
		return
	}
 
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(s.Alpha)
	op.GeoM.Scale(scaleX, scaleY)

	op.GeoM.Translate(x, y)
	screen.DrawImage(s.Image, op)
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
