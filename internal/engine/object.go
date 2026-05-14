package engine

import "errors"

type Object struct {
	ID     string
	TileX  int
	TileY  int
	Layer  int
	SizeX  int
	SizeY  int
	ClusterMultiplier float32
	Sprite *Sprite
}

func NewObject(id string, tileX, tileY int, layer int, sizeX, sizeY int, clusterMultiplier float32, sprite *Sprite) *Object {
	return &Object{
		ID:     id,
		TileX:  tileX,
		TileY:  tileY,
		Layer:  layer,
		SizeX:  sizeX,
		SizeY:  sizeY,
		ClusterMultiplier: clusterMultiplier,
		Sprite: sprite,
	}
}

func (o *Object) Clone() *Object {
	newObject := *o
	newObject.Sprite = o.Sprite.Instantiate()

	return &newObject
}

func (m *Map) PutObject(object *Object) ( error) {

	if object.TileX+object.SizeX > m.MapSizeX || object.TileY+object.SizeY > m.MapSizeY {
		return errors.New("objeto out of bounds")
	}

	for y := object.TileY; y < object.TileY+object.SizeY; y++ {
		for x := object.TileX; x < object.TileX+object.SizeX; x++ {
			if object.Layer == 0 && m.TilesZ0[y][x] != nil {
				return errors.New("objeto conflitante com outro no mapa")
			} else if object.Layer == 1 && m.TilesZ1[y][x] != nil {
				return errors.New("objeto conflitante com outro no mapa")
			} else if object.Layer == 2 && m.TilesZ2[y][x] != nil {
				return errors.New("objeto conflitante com outro no mapa")
			}
		}
	}
	for y := object.TileY; y < object.TileY+object.SizeY; y++ {
		for x := object.TileX; x < object.TileX+object.SizeX; x++ {
			switch object.Layer {
				case 0:
					m.TilesZ0[y][x] = object
				case 1:
					m.TilesZ1[y][x] = object
				case 2:
					m.TilesZ2[y][x] = object
			}
		}
	}

	return nil
}

func (m *Map) DestroyObject(object *Object) {
	for y := object.TileY; y < object.TileY+object.SizeY; y++ {
		for x := object.TileX; x < object.TileX+object.SizeX; x++ {
			switch object.Layer {
				case 0:
					m.TilesZ0[y][x] = nil
				case 1:
					m.TilesZ1[y][x] = nil
				case 2:
					m.TilesZ2[y][x] = nil
			}
		}
	}
}

