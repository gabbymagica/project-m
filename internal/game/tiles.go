package game

import (
	"fmt"
	"image/color"
	"math"
	"project_m/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) HandleTileDrawStates(screen *ebiten.Image) {
	if g.TileStates[StatePreview] {
		g.DrawTileBorders(screen)
	}
	if g.TileStates[StateBuilding] {
		g.DrawSpritePreview(screen, g.TestObject)
	}
	if g.TileStates[StateDestroying] {
		g.DrawDestroyingPreview(screen)
	}
	if g.TileStates[StateDrawSprites] {
		g.DrawTileSprites(screen)
	}
}

func (g *Game) DrawTileBorders(screen *ebiten.Image) {
	bounds := screen.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	tile_x0 := float64(g.Camera.X - (width / 2))
	tile_y0 := float64(g.Camera.Y - (height / 2))
	tileSize := float64(g.Map.Tilesize)

	minTileX := int(math.Floor(tile_x0 / tileSize))
	maxTileX := int(math.Floor((tile_x0+float64(width))/tileSize)) + 1
	minTileY := int(math.Floor(tile_y0 / tileSize))
	maxTileY := int(math.Floor((tile_y0+float64(height))/tileSize)) + 1

	for tileY := minTileY; tileY <= maxTileY; tileY++ {
		if tileY < 0 || tileY >= g.Map.MapSizeY {
			continue
		}
		for tileX := minTileX; tileX <= maxTileX; tileX++ {
			if tileX < 0 || tileX >= g.Map.MapSizeX {
				continue
			}

			screenX := float32((float64(tileX) * tileSize) - tile_x0)
			screenY := float32((float64(tileY) * tileSize) - tile_y0)
			ts := float32(g.Map.Tilesize)

			vector.StrokeLine(screen, screenX, screenY, screenX+ts, screenY, 3, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			vector.StrokeLine(screen, screenX, screenY, screenX, screenY+ts, 3, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)

			if tileX == g.Map.MapSizeX-1 {
				vector.StrokeLine(screen, screenX+ts, screenY, screenX+ts, screenY+ts, 3, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			}
			if tileY == g.Map.MapSizeY-1 {
				vector.StrokeLine(screen, screenX, screenY+ts, screenX+ts, screenY+ts, 3, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			}
		}
	}
}

func (g *Game) DrawTileSprites(screen *ebiten.Image) {
	bounds := screen.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	tile_x0 := float64(g.Camera.X - (width / 2))
	tile_y0 := float64(g.Camera.Y - (height / 2))
	tileSize := float64(g.Map.Tilesize)

	minTileX := int(math.Floor(tile_x0 / tileSize))
	maxTileX := int(math.Floor((tile_x0+float64(width))/tileSize)) + 1
	minTileY := int(math.Floor(tile_y0 / tileSize))
	maxTileY := int(math.Floor((tile_y0+float64(height))/tileSize)) + 1

	for z := 0; z < 3; z++ {
		for tileY := minTileY; tileY <= maxTileY; tileY++ {
			if tileY < 0 || tileY >= g.Map.MapSizeY {
				continue
			}
			for tileX := minTileX; tileX <= maxTileX; tileX++ {
				if tileX < 0 || tileX >= g.Map.MapSizeX {
					continue
				}

				var object *engine.Object
				if z == 0 {
					object = g.Map.TilesZ0[tileY][tileX]
				} else if z == 1 {
					object = g.Map.TilesZ1[tileY][tileX]
				} else if z == 2 {
					object = g.Map.TilesZ2[tileY][tileX]
				}

				if object != nil && object.Sprite != nil {
					if object.TileX == tileX && object.TileY == tileY {
						screenX := (float64(tileX) * tileSize) - tile_x0
						screenY := (float64(tileY) * tileSize) - tile_y0

						object.Sprite.X = int(screenX)
						object.Sprite.Y = int(screenY)
						object.Sprite.ScaleX = (tileSize * float64(object.SizeX)) / 16
						object.Sprite.ScaleY = (tileSize * float64(object.SizeY)) / 16
						object.Sprite.Draw(screen)
					}
				}
			}
		}
	}
}

func (g *Game) DrawSpritePreview(screen *ebiten.Image, object *engine.Object) {
	x, y := ebiten.CursorPosition()
	bounds := screen.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	tileSize := float64(g.Map.Tilesize)

	tile_x0 := float64(g.Camera.X) - (float64(width) / 2)
	tile_y0 := float64(g.Camera.Y) - (float64(height) / 2)

	worldCursorX := tile_x0 + float64(x)
	worldCursorY := tile_y0 + float64(y)

	tileX := int(math.Floor(worldCursorX / tileSize))
	tileY := int(math.Floor(worldCursorY / tileSize))

	screenX := int((float64(tileX) * tileSize) - tile_x0)
	screenY := int((float64(tileY) * tileSize) - tile_y0)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\nCursor X: %d, Cursor Y: %d\nScreenX: %d, ScreenY: %d", x, y, screenX, screenY))

	object.Sprite.X = screenX
	object.Sprite.Y = screenY
	object.Sprite.ScaleX = (tileSize * float64(object.SizeX)) / 16
	object.Sprite.ScaleY = (tileSize * float64(object.SizeY)) / 16
	object.Sprite.SetAlpha(0.4)
	object.Sprite.Draw(screen)
}

func (g *Game) DrawDestroyingPreview(screen *ebiten.Image) {
	x, y := ebiten.CursorPosition()
	bounds := screen.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	tileSize := float64(g.Map.Tilesize)

	tile_x0 := float64(g.Camera.X) - (float64(width) / 2)
	tile_y0 := float64(g.Camera.Y) - (float64(height) / 2)

	worldCursorX := tile_x0 + float64(x)
	worldCursorY := tile_y0 + float64(y)

	tileX := int(math.Floor(worldCursorX / tileSize))
	tileY := int(math.Floor(worldCursorY / tileSize))

	if tileX < 0 || tileX >= g.Map.MapSizeX || tileY < 0 || tileY >= g.Map.MapSizeY {
		return
	}

	var object *engine.Object
	if g.Map.TilesZ0[tileY][tileX] != nil {
		object = g.Map.TilesZ0[tileY][tileX]
	}
	if g.Map.TilesZ1[tileY][tileX] != nil {
		object = g.Map.TilesZ1[tileY][tileX]
	}
	if g.Map.TilesZ2[tileY][tileX] != nil {
		object = g.Map.TilesZ1[tileY][tileX]
	}
	if object == nil {
		return
	}

	screenX := float32((float64(object.TileX) * tileSize) - tile_x0)
	screenY := float32((float64(object.TileY) * tileSize) - tile_y0)
	ts := float32(g.Map.Tilesize)

	red := color.RGBA{0x88, 0x00, 0x00, 0xFF}

	sX := float32(screenX)
	sY := float32(screenY)
	rectWidth := float32(ts) * float32(object.SizeX)
	rectHeight := float32(ts) * float32(object.SizeY)

	vector.StrokeRect(screen, sX, sY, rectWidth, rectHeight, 3, red, false)
}
