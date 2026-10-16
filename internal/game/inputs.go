package game

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) HandleGameInputs() {
	g.inputMovePlayer()
	g.inputUIStates()
	g.inputMouse()
}

func (g *Game) inputMouse() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()
		
		tileX, tileY := g.Camera.ScreenToTile(float64(posX), float64(posY), g.Map.Tilesize)

		if (tileX > 0 && tileX < g.Map.MapSizeX) && (tileY > 0 && tileY < g.Map.MapSizeY) {
			tile := g.Map.Tiles[tileY][tileX]

			if tile != nil {
				if g.Player.IsMinerable(tile.ID) {
					g.Player.StartMining(tile, tileX, tileY)
				}
			}
		}	
	}
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
		g.Player.Move(g.Player.Speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.Move(-g.Player.Speed, 0)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Move(0, g.Player.Speed)
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Move(0, -g.Player.Speed)
	}
}
