package engine

type Map struct {
	TilesZ0            [32][100]*Object
	TilesZ1            [32][100]*Object
	TilesZ2            [32][100]*Object
	MapSizeX, MapSizeY int
	Tilesize           int
}

func NewMap(tilesize int) *Map {
	return &Map{
		TilesZ0:  [32][100]*Object{},
		TilesZ1:  [32][100]*Object{},
		TilesZ2:  [32][100]*Object{},
		MapSizeY: 32,
		MapSizeX: 100,
		Tilesize: tilesize,
	}
}
