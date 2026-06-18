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
		g.GameMode = ModeBuilding
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		g.GameMode = ModeDestroying
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.RenderFlags.ShowTileBorders = !g.RenderFlags.ShowTileBorders
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
