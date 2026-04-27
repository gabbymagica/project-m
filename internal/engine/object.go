package engine

import "errors"

type Object struct {
	ID     int
	tileX  int
	tileY  int
	sizeX  int
	sizeY  int
	sprite *Sprite
}

func (m *Map) NewObject(tileX, tileY int, sizeX, sizeY int, sprite *Sprite) (*Object, error) {
	object := &Object{
		tileX:  tileX,
		tileY:  tileY,
		sizeX:  sizeX,
		sizeY:  sizeY,
		sprite: sprite,
	}

	if tileX+sizeX >= m.MapSizeX || tileY+sizeY >= m.MapSizeY {
		return nil, errors.New("objeto out of bounds")
	}

	// verifica se tem algum objeto no lugar
	for y := tileY; y < tileY+sizeY; y++ {
		for x := tileX; x < tileX+sizeX; x++ {
			if m.tiles[x][y] != nil {
				return nil, errors.New("objeto conflitante com outro no mapa")
			}
		}
	}

	for y := tileY; y < tileY+sizeY; y++ {
		for x := tileX; x < tileX+sizeX; x++ {
			m.tiles[x][y] = object
		}
	}

	return object, nil
}
