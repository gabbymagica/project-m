package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) HandleGameInputs() {
	g.inputMovePlayer()
	g.inputUIStates()
}

func (g *Game) inputUIStates() {
	if inpututil.IsKeyJustPressed(ebiten.KeyB) {
		g.TileStates[StateBuilding] = !g.TileStates[StateBuilding]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.TileStates[StateDestroying] = !g.TileStates[StateDestroying]
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.TileStates[StatePreview] = !g.TileStates[StatePreview]
	}
}

func (g *Game) inputMovePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.X += 5
		g.Camera.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.X -= 5
		g.Camera.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Y += 5
		g.Camera.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Y -= 5
		g.Camera.Y -= 5
	}
}
