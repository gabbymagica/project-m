package ui

import (
	"project_m/internal/engine"

	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type GameFonts struct {
	InventoryUIFont *text.GoTextFaceSource
	ToolTipFont     *text.GoTextFaceSource
}

func InitFonts() *GameFonts {
	return &GameFonts{
		InventoryUIFont: engine.LoadFont("fonts/Pix32.ttf"),
		ToolTipFont:     engine.LoadFont("fonts/Arial.ttf"),
	}
}
