package game

type GameMode int

const (
	ModeNormal GameMode = iota
	ModeBuilding
	ModeDestroying
)

type RenderFlags struct {
	ShowTileBorders bool
	ShowSprites     bool
}
