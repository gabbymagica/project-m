package engine

import (
	"project_m/internal/constants"

	"github.com/hajimehoshi/ebiten/v2"
)

type BuildInterface interface {
	Update(dt float64)
	Draw(screen *ebiten.Image, x, y float64)
}

type Build struct {
	ID           	constants.BuildID
	TileX, TileY 	int
	Solid 	  		bool
	SizeX, SizeY 	int		// Quantidade de tiles que ocupa a construção
	Sprite 			*Sprite
}
