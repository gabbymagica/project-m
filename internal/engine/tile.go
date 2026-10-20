package engine

import (
	"project_m/internal/constants"
)

type Tile struct {
	ID     constants.TileID
	Sprite *Sprite
	Solid  bool
}

func NewTile(id constants.TileID, solid bool, sprite *Sprite) *Tile {
	return &Tile{
		ID:     id,
		Sprite: sprite,
		Solid:  solid,
	}
}

type AutoGenerateTile struct {
	Tile              *Tile
	ClusterMultiplier float32
}

func NewAutoGenerateTile(tile *Tile, clusterMultiplier float32) *AutoGenerateTile {
	return &AutoGenerateTile{
		Tile:              tile,
		ClusterMultiplier: clusterMultiplier,
	}
}
