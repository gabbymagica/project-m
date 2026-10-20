package game

import (
	"image/color"
	"project_m/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) HandleRenderFlags(screen *ebiten.Image) {
	if g.RenderFlags.ShowSprites {
		g.DrawTileSprites(screen)
	}
	if g.RenderFlags.ShowBuildPreview {
		if g.Player.Building.SelectedBuildType != nil {
			g.DrawSpritePreview(screen, g.Player.Building.SelectedBuildType)
		}
	}
	if g.RenderFlags.ShowDestroyPreview {
		g.DrawDestroyingPreview(screen)
	}
	if g.RenderFlags.ShowTileBorders {
		g.DrawTileBorders(screen)
	}
}

func (g *Game) DrawTileBorders(screen *ebiten.Image) {
	minTileX, minTileY, maxTileX, maxTileY := g.Camera.VisibleTileBounds(
		g.Map.Tilesize,
		g.Map.MapSizeX,
		g.Map.MapSizeY,
	)

	ts := float64(g.Map.Tilesize)
	gridColor := color.RGBA{0xFF, 0xFF, 0xFF, 0x40}
	lineThickness := float32(1.0)

	screenMinX, screenMinY := g.Camera.WorldToScreen(float64(minTileX)*ts, float64(minTileY)*ts)
	screenMaxX, screenMaxY := g.Camera.WorldToScreen(float64(maxTileX)*ts, float64(maxTileY)*ts)

	for tx := minTileX; tx <= maxTileX; tx++ {
		sx, _ := g.Camera.WorldToScreen(float64(tx)*ts, 0)
		vector.StrokeLine(
			screen,
			float32(sx), float32(screenMinY),
			float32(sx), float32(screenMaxY),
			lineThickness,
			gridColor,
			false,
		)
	}

	for ty := minTileY; ty <= maxTileY; ty++ {
		_, sy := g.Camera.WorldToScreen(0, float64(ty)*ts)
		vector.StrokeLine(
			screen,
			float32(screenMinX), float32(sy),
			float32(screenMaxX), float32(sy),
			lineThickness,
			gridColor,
			false,
		)
	}
}

func (g *Game) DrawTileSprites(screen *ebiten.Image) {
	minX, minY, maxX, maxY := g.Camera.VisibleTileBounds(g.Map.Tilesize, g.Map.MapSizeX, g.Map.MapSizeY)

	for tileY := minY; tileY < maxY; tileY++ {
		for tileX := minX; tileX < maxX; tileX++ {
			tile := g.Map.Tiles[tileY][tileX]
			if tile == nil || tile.Sprite == nil {
				continue
			}

			worldX := float64(tileX * g.Map.Tilesize)
			worldY := float64(tileY * g.Map.Tilesize)
			screenX, screenY := g.Camera.WorldToScreen(worldX, worldY)

			tile.Sprite.Draw(screen, screenX, screenY, 1, 1)

		}
	}

}

func (g *Game) DrawSpritePreview(screen *ebiten.Image, build *engine.BuildType) {
	if build.Sprite == nil {
		return
	}

	cx, cy := ebiten.CursorPosition()
	tileX, tileY := g.Camera.ScreenToTile(float64(cx), float64(cy), g.Map.Tilesize)

	if tileX < 0 || tileY < 0 || tileX+build.SizeX > g.Map.MapSizeX || tileY+build.SizeY > g.Map.MapSizeY {
		return
	}

	worldX := float64(tileX * g.Map.Tilesize)
	worldY := float64(tileY * g.Map.Tilesize)
	screenX, screenY := g.Camera.WorldToScreen(worldX, worldY)

	oldAlpha := build.Sprite.Alpha
	build.Sprite.SetAlpha(0.4)
	build.Sprite.Draw(screen, screenX, screenY, 1.0, 1.0)
	build.Sprite.SetAlpha(oldAlpha)
}

func (g *Game) DrawDestroyingPreview(screen *ebiten.Image) {
	cx, cy := ebiten.CursorPosition()
	tileX, tileY := g.Camera.ScreenToTile(float64(cx), float64(cy), g.Map.Tilesize)

	if tileX < 0 || tileX >= g.Map.MapSizeX || tileY < 0 || tileY >= g.Map.MapSizeY {
		return
	}

	var target *engine.Build
	if obj := g.Map.Builds[tileY][tileX]; obj != nil {
		target = obj
	}

	if target == nil {
		return
	}

	worldX := float64(tileX * g.Map.Tilesize)
	worldY := float64(tileY * g.Map.Tilesize)
	screenX, screenY := g.Camera.WorldToScreen(worldX, worldY)

	ts := float32(g.Map.Tilesize)
	rectWidth := ts * float32(target.BuildType.SizeX)
	rectHeight := ts * float32(target.BuildType.SizeY)
	sX := float32(screenX)
	sY := float32(screenY)

	fillRed := color.RGBA{0xFF, 0x00, 0x00, 0x44}
	strokeRed := color.RGBA{0xFF, 0x22, 0x22, 0xEE}

	vector.FillRect(screen, sX, sY, rectWidth, rectHeight, fillRed, false)
	vector.StrokeRect(screen, sX, sY, rectWidth, rectHeight, 2, strokeRed, false)
}
