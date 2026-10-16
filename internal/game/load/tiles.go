package load

import (
	"project_m/internal/constants"
	"project_m/internal/engine"
)

var TILES = map[constants.TileID]*engine.Tile{
	constants.TileGrass:     engine.NewTile(constants.TileGrass, false, TILE_SPRITES[constants.TileGrass]),
	constants.TileOreCopper: engine.NewTile(constants.TileOreCopper, true, TILE_SPRITES[constants.TileOreCopper]),
}