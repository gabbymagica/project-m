package engine

type Map struct {
	tiles              [32][100]*Object
	MapSizeX, MapSizeY int
	Tilesize           int
}

func NewMap(tilesize int) *Map {
	return &Map{
		tiles:    [32][100]*Object{},
		MapSizeY: 32,
		MapSizeX: 100,
		Tilesize: tilesize,
	}
}
