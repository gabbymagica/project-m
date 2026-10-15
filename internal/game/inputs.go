package game

import (

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"project_m/internal/constants"
)

func (g *Game) HandleGameInputs() {
	g.inputMovePlayer()
	g.inputUIStates()
	g.inputMouse()
}

func (g *Game) inputMouse() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()
		posX += int(g.Camera.X - constants.LAYOUT_WIDTH/2)
		posY += int(g.Camera.Y - constants.LAYOUT_HEIGHT/2)

		tileX := posX / constants.TILESIZE
		tileY := posY / constants.TILESIZE

		if (tileX > 0 && tileX < g.Map.MapSizeX) && (tileY > 0 && tileY < g.Map.MapSizeY) {
			verifyLayers:
			for layer := constants.MAP_LAYERS-1; layer >= 0; layer-- {
				object := g.Map.Tiles[layer][tileY][tileX]

				if object != nil {
					switch layer {
					case constants.LAYER_GROUND:
						if g.Player.IsMinerable(object.ID) {
							g.Player.StartMining(object)
						}
						break verifyLayers
					case constants.LAYER_OBJECTS:
						break verifyLayers
					case constants.LAYER_ENTITIES:
						break verifyLayers
					}
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
