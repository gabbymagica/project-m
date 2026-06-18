package engine

import (
	"math/rand"
)

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

func (m *Map) GenerateMap(priorities map[*Object]float32) {
	posicoes := [][2]int{
		{0, -1},
		{-1, 0},
		{-1, -1},
		{1, -1},
	}

	prioridadesLocais := make(map[*Object]float32)

	for y := 0; y < m.MapSizeY; y++ {
		for x := 0; x < m.MapSizeX; x++ {

			clear(prioridadesLocais)

			var prioridadeTotal float32 = 0.0

			for object, priority := range priorities {
				prioridadesLocais[object] = priority

				for _, p := range posicoes {
					ny := y + p[0]
					nx := x + p[1]

					if ny < 0 || ny >= m.MapSizeY || nx < 0 || nx >= m.MapSizeX || m.TilesZ0[ny][nx] == nil {
						continue
					}

					object2 := m.TilesZ0[ny][nx]

					if object2.ID == object.ID {
						prioridadesLocais[object] += object.ClusterMultiplier
					}
				}

				prioridadeTotal += prioridadesLocais[object]
			}

			i := float32(rand.Intn(int(prioridadeTotal)) + 1)

			for object, priority := range prioridadesLocais {
				i -= priority

				if i <= 0 {

					object := object.Clone()
					object.TileX = x
					object.TileY = y

					m.PutObject(object)

					break
				}
			}
		}
	}
}
