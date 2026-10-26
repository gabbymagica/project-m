package load

import (
	"image/color"
	"project_m/assets"
	"project_m/internal/constants"
	"project_m/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func el_quadrado_generator(clr [3]uint8, sizeX, sizeY int) *ebiten.Image {
	img := ebiten.NewImage(sizeX*32, sizeY*32)
	img.Fill(color.RGBA{clr[0], clr[1], clr[2], 0xff})
	return img
}

var BUILD_SPRITES = map[constants.BuildID]*engine.Sprite{
	constants.BuildTest1: engine.NewSprite("TEST1", "Test1 Build", el_quadrado_generator([3]uint8{255, 0, 0}, 2, 2)),
	constants.BuildTest2: engine.NewSprite("TEST1", "Test2 Build", el_quadrado_generator([3]uint8{0, 255, 0}, 1, 1)),
}

var TILE_SPRITES = map[constants.TileID]*engine.Sprite{
	constants.TileGrass:     engine.NewSprite("GRASS", "Grass", assets.LoadImage("tiles/grass.png")),
	constants.TileOreCopper: engine.NewSprite("ORE_COPPER", "Copper Ore", assets.LoadImage("tiles/ore_copper.png")),
}

var PLAYER_SPRITES = map[string]*engine.Sprite{
	"main": engine.NewSprite("PLAYER", "Player", el_quadrado_generator([3]uint8{0, 0, 0}, 1, 1)),
}
