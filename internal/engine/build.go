package engine

import (
	"project_m/internal/constants"
)

type BuildType struct {
	ID                  constants.BuildID
	Solid               bool
	SizeX, SizeY        int
	Sprite              *Sprite
	BuildingTimeSeconds float64
}

type Build struct {
	BuildType *BuildType // <- ponteiro de referência pro tipo. Permite a edição, é só não ser bobão e editar xd

	TileX, TileY     int
	BuildProgress    float64
	FinishedBuilding bool
}

func NewBuildType(id constants.BuildID, solid bool, sizeX, sizeY int, sprite *Sprite, bts float64) *BuildType {
	return &BuildType{
		ID:                  id,
		Solid:               solid,
		SizeX:               sizeX,
		SizeY:               sizeY,
		Sprite:              sprite,
		BuildingTimeSeconds: bts,
	}
}
