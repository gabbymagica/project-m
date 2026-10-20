package game

import (
	"project_m/internal/constants"
	"project_m/internal/game/load"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) HandleGameInputs() {
	g.inputPlayer()
	g.inputRenderFlags()
	g.inputMouse()
}

func (g *Game) inputMouse() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		posX, posY := ebiten.CursorPosition()
		tileX, tileY := g.Camera.ScreenToTile(float64(posX), float64(posY), g.Map.Tilesize)

		if (tileX <= 0 || tileX >= g.Map.MapSizeX) || (tileY <= 0 || tileY >= g.Map.MapSizeY) {
			return
		}

		switch g.Player.Action {
		case constants.PreviewBuilding:
			if g.Player.CanBuild(tileX, tileY) && g.Player.Building.SelectedBuildType != nil {
				var build = g.Map.PutBuilding(g.Player.Building.SelectedBuildType, tileX, tileY)
				g.Player.StartBuilding(build)
			}
			return
		case constants.Idling:
			clickBuild := g.Map.Builds[tileY][tileX]
			if clickBuild.FinishedBuilding == false {
				g.Player.StartBuilding(clickBuild)
				return
			}

			tile := g.Map.Tiles[tileY][tileX]
			if tile != nil {
				if g.Player.IsMinerable(tile.ID) {
					if g.Player.CanMine(tileX, tileY) {
						g.Player.StartMining(tile, tileX, tileY)
					}
				}
			}
			return
		}
	}
}

func (g *Game) inputRenderFlags() {
	if inpututil.IsKeyJustPressed(ebiten.KeyT) {
		g.RenderFlags.ShowTileBorders = !g.RenderFlags.ShowTileBorders
	}
}

func (g *Game) inputPlayer() {
	// movement
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

	// actions
	if ebiten.IsKeyPressed(ebiten.KeyB) {
		if g.Player.Action == constants.Idling {
			g.Player.Action = constants.PreviewBuilding
			g.RenderFlags.ShowBuildPreview = true
		} else if g.Player.Action == constants.PreviewBuilding {
			g.Player.Action = constants.Idling
			g.RenderFlags.ShowBuildPreview = false
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		if g.Player.Action == constants.Idling {
			g.Player.Action = constants.PreviewDestroying
			g.RenderFlags.ShowDestroyPreview = true
		} else if g.Player.Action == constants.PreviewDestroying {
			g.Player.Action = constants.Idling
			g.RenderFlags.ShowDestroyPreview = false
		}
	}

	// debug
	if ebiten.IsKeyPressed(ebiten.Key1) {
		if g.Player.Action != constants.Building {
			g.Player.Building.SelectedBuildType = load.BUILD_TYPES[constants.BuildTest1]
		}
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		if g.Player.Action != constants.Building {
			g.Player.Building.SelectedBuildType = load.BUILD_TYPES[constants.BuildTest2]
		}
	}
}
