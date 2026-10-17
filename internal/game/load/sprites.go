package load

import (
	"image/color"

	"project_m/assets"
	"project_m/internal/constants"
	"project_m/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

var TILE_SPRITES = map[constants.TileID]*engine.Sprite{
	constants.TileGrass : engine.NewSprite("GRASS", "Grass", assets.LoadImage("tiles/grass.png")),
	constants.TileOreCopper : engine.NewSprite("ORE_COPPER", "Copper Ore", assets.LoadImage("tiles/ore_copper.png")),
}

var el_quadrado_preto = func() *ebiten.Image {
	img := ebiten.NewImage(32, 32)
	img.Fill(color.RGBA{0, 0, 0, 0xff})
	return img
}()
var PLAYER_SPRITES = map[string]*engine.Sprite{
	"main": engine.NewSprite("PLAYER", "Player", el_quadrado_preto),
}