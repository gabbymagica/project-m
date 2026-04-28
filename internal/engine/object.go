package engine

import "errors"

type Object struct {
	ID     int
	TileX  int
	TileY  int
	SizeX  int
	SizeY  int
	Sprite *Sprite
}

func (m *Map) NewObject(tileX, tileY int, tileZ int, sizeX, sizeY int, sprite *Sprite) (*Object, error) {
	object := &Object{
		TileX:  tileX,
		TileY:  tileY,
		SizeX:  sizeX,
		SizeY:  sizeY,
		Sprite: sprite,
	}

	if tileX+sizeX >= m.MapSizeX || tileY+sizeY >= m.MapSizeY {
		return nil, errors.New("objeto out of bounds")
	}

	// verifica se tem algum objeto no lugar
	for y := tileY; y < tileY+sizeY; y++ {
		for x := tileX; x < tileX+sizeX; x++ {
			if tileZ == 0 {
				if m.TilesZ0[y][x] != nil {
					return nil, errors.New("objeto conflitante com outro no mapa")
				}
			} else if tileZ == 1 {
				if m.TilesZ1[y][x] != nil {
					return nil, errors.New("objeto conflitante com outro no mapa")
				}
			} else if tileZ == 2 {
				if m.TilesZ2[y][x] != nil {
					return nil, errors.New("objeto conflitante com outro no mapa")
				}
			}
		}
	}

	for y := tileY; y < tileY+sizeY; y++ {
		for x := tileX; x < tileX+sizeX; x++ {
			if tileZ == 0 {
				m.TilesZ0[y][x] = object
			} else if tileZ == 1 {
				m.TilesZ1[y][x] = object
			} else if tileZ == 2 {
				m.TilesZ2[y][x] = object
			}
		}
	}

	return object, nil
}
