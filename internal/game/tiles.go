package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) DrawTileBorders(screen *ebiten.Image) {
	bounds := screen.Bounds()
	height := bounds.Dy()
	width := bounds.Dx()

	tile_x0, tile_y0 := g.Player.X-(width/2), g.Player.Y-(height/2)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\ntile_x0: %d tile_y0: %d", tile_x0, tile_y0))

	for y := tile_y0; y <= tile_y0+height+16; y += g.Map.Tilesize {
		for x := tile_x0; x <= tile_x0+width+16; x += g.Map.Tilesize {
			tileX, tileY := int(x/g.Map.Tilesize), int(y/g.Map.Tilesize)
			if tileY < 0 || tileY >= g.Map.MapSizeY {
				continue
			}
			if tileX < 0 || tileX >= g.Map.MapSizeX {
				continue
			}
			tile_edgeX, tile_edgeY := x-(x%g.Map.Tilesize), y-(y%g.Map.Tilesize)
			screen_tile_edgeX, screen_tile_edgeY := tile_edgeX-tile_x0, tile_edgeY-tile_y0
			vector.StrokeLine(screen, float32(screen_tile_edgeX), float32(screen_tile_edgeY), float32(screen_tile_edgeX)+float32(g.Map.Tilesize), float32(screen_tile_edgeY), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			vector.StrokeLine(screen, float32(screen_tile_edgeX), float32(screen_tile_edgeY), float32(screen_tile_edgeX), float32(screen_tile_edgeY)+float32(g.Map.Tilesize), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			if tileX == g.Map.MapSizeX-1 {
				vector.StrokeLine(screen, float32(screen_tile_edgeX)+float32(g.Map.Tilesize), float32(screen_tile_edgeY), float32(screen_tile_edgeX)+float32(g.Map.Tilesize), float32(screen_tile_edgeY)+float32(g.Map.Tilesize), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			}
			if tileY == g.Map.MapSizeY-1 {
				vector.StrokeLine(screen, float32(screen_tile_edgeX), float32(screen_tile_edgeY)+float32(g.Map.Tilesize), float32(screen_tile_edgeX)+float32(g.Map.Tilesize), float32(screen_tile_edgeY)+float32(g.Map.Tilesize), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			}
		}
	}
}

func (g *Game) DrawTileSprites(screen *ebiten.Image) {

}
