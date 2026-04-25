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

	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n\nheight: %d\nwidth: %d", height, width))

	tileSize := g.Tilesize

	x0, y0 := g.Player.X-float64(width)/2, g.Player.Y-float64(height)/2
	tileY0 := int(y0 / float64(tileSize))
	tileX0 := int(x0 / float64(tileSize))
	maxTilesHeight := int(height / tileSize)
	maxTilesWidth := int(width / tileSize)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("\n\n\n\n\nx0: %d\ny0: %d\n tilex0: %d\n tiley0: %d\n", x0, y0, tileX0, tileY0))

	for y := tileY0; y <= tileY0+maxTilesHeight; y++ {
		if y < 0 || y > len(g.Map) {
			continue
		}
		for x := tileX0; x <= tileX0+maxTilesWidth; x++ {
			if x < 0 || x > len(g.Map[y]) {
				continue
			}
			tileXPos := float32((x - tileX0) * tileSize)
			tileYPos := float32((y - tileY0) * tileSize)
			vector.StrokeLine(screen, tileXPos, tileYPos, tileXPos+float32(g.Tilesize), tileYPos, 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			vector.StrokeLine(screen, tileXPos, tileYPos, tileXPos, tileYPos+float32(g.Tilesize), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			if x == len(g.Map[y]) {
				vector.StrokeLine(screen, tileXPos+float32(g.Tilesize), tileYPos, tileXPos+float32(g.Tilesize), tileYPos+float32(g.Tilesize), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			}
			if y == len(g.Map) {
				vector.StrokeLine(screen, tileXPos, tileYPos+float32(g.Tilesize), tileXPos+float32(g.Tilesize), tileYPos+float32(g.Tilesize), 1, color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, false)
			}
		}
	}
}
