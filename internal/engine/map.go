package engine

import (
	"math/rand"
)

type Map struct {
	Tiles [][]*Tile // Camada de Tiles do chão

	Builds     [][]*Build // Camada de builds para otimização de busca.
	BuildsList []*Build   // Lista de builds para otimização de renderização.

	MapSizeX, MapSizeY int
	Tilesize           int
}

func NewMap(tilesize int, sizeX, sizeY int) *Map {
	tiles := make([][]*Tile, sizeY)
	build := make([][]*Build, sizeY)

	for y := 0; y < sizeY; y++ {
		tiles[y] = make([]*Tile, sizeX)
		build[y] = make([]*Build, sizeX)
	}

	return &Map{
		Tiles:    tiles,
		Builds:   build,
		MapSizeY: sizeY,
		MapSizeX: sizeX,
		Tilesize: tilesize,
	}
}

func (m *Map) GenerateMap(priorities map[*AutoGenerateTile]float32) {
	posicoes := [][2]int{
		{0, -1},
		{-1, 0},
		{-1, -1},
		{-1, 1},
	}

	prioridadesLocais := make(map[*AutoGenerateTile]float32)

	for y := 0; y < m.MapSizeY; y++ {
		for x := 0; x < m.MapSizeX; x++ {

			clear(prioridadesLocais)

			var prioridadeTotal float32 = 0.0

			for tile, priority := range priorities {
				prioridadesLocais[tile] = priority

				for _, p := range posicoes {
					ny := y + p[0]
					nx := x + p[1]

					if ny < 0 || ny >= m.MapSizeY || nx < 0 || nx >= m.MapSizeX || m.Tiles[ny][nx] == nil {
						continue
					}

					tile2 := m.Tiles[ny][nx]

					if tile2.ID == tile.Tile.ID {
						prioridadesLocais[tile] += tile.ClusterMultiplier
					}
				}

				prioridadeTotal += prioridadesLocais[tile]
			}

			i := float32(rand.Intn(int(prioridadeTotal)) + 1)

			for tile, priority := range prioridadesLocais {
				i -= priority

				if i <= 0 {

					m.PutTile(tile.Tile, x, y)

					break
				}
			}
		}
	}
}

func (m *Map) Update() {
	// verifica por construções com flag destroyed = true para serem removidas
	for i := len(m.BuildsList) - 1; i >= 0; i-- {
		v := m.BuildsList[i]
		if v.Destroyed {
			m.BuildsList[i] = m.BuildsList[len(m.BuildsList)-1]
			m.BuildsList[len(m.BuildsList)-1] = nil
			m.BuildsList = m.BuildsList[:len(m.BuildsList)-1]

			for y := v.TileY; y < v.TileY+v.BuildType.SizeY; y++ {
				for x := v.TileX; x < v.TileX+v.BuildType.SizeX; x++ {
					m.Builds[y][x] = nil
				}
			}
		}
	}
}

func (m *Map) PutTile(tile *Tile, x, y int) {
	if x < 0 || x >= m.MapSizeX || y < 0 || y >= m.MapSizeY {
		return
	}
	m.Tiles[y][x] = tile
}

func (m *Map) PutBuilding(buildType *BuildType, x, y int) *Build {
	if x < 0 || x+buildType.SizeX > m.MapSizeX || y < 0 || y+buildType.SizeY > m.MapSizeY {
		return nil
	}

	// verificacao
	for iplison := y; iplison < y+buildType.SizeY; iplison++ {
		for xis := x; xis < x+buildType.SizeX; xis++ {
			if m.Builds[iplison][xis] != nil {
				return nil
			}
		}
	}

	buildInstance := &Build{
		BuildType:        buildType,
		TileX:            x,
		TileY:            y,
		BuildProgress:    0,
		FinishedBuilding: false,
	}

	// colocacao ..?
	for iplison := y; iplison < y+buildType.SizeY; iplison++ {
		for xis := x; xis < x+buildType.SizeX; xis++ {
			m.Builds[iplison][xis] = buildInstance
		}
	}

	m.BuildsList = append(m.BuildsList, buildInstance)
	return buildInstance
}
