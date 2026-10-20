package player

import (
	"image/color"
	"project_m/internal/constants"
	"project_m/internal/engine"
	item "project_m/internal/game/items"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type TargetTile struct {
	ID           constants.TileID
	TileX, TileY int
}

type Mining struct {
	TargetTile    *TargetTile
	Range         float64
	MiningDelay   float64
	Timer         float64
	MinerableOres []constants.TileID
}

type Building struct {
	SelectedBuildType *engine.BuildType
	BuildTarget       *engine.Build
	Range             float64
}

type Player struct {
	engine.Entity
	Inventory *inventory.Inventory
	Action    constants.ActionID
	Mining    *Mining
	Building  *Building
}

func (p *Player) Update(dt float64) {
	p.UpdateMining(dt)
}

func (p *Player) UpdateMining(dt float64) {
	if p.Action == constants.Mining && p.Mining.TargetTile != nil {
		objectX := float64(p.Mining.TargetTile.TileX*constants.TILESIZE + constants.TILESIZE/2)
		objectY := float64(p.Mining.TargetTile.TileY*constants.TILESIZE + constants.TILESIZE/2)
		if engine.Distance(p.X, p.Y, objectX, objectY) > p.Mining.Range {
			p.StopMining()
			return
		}

		p.Mining.Timer += dt
		if p.Mining.Timer >= p.Mining.MiningDelay {
			p.Mining.Timer = 0
			p.Inventory.AddItem(item.Manager.Items[constants.ItemID(p.Mining.TargetTile.ID)], 1)
		}

	}
}

func (p *Player) IsMinerable(ID constants.TileID) bool {
	for _, ore := range p.Mining.MinerableOres {
		if ore == ID {
			return true
		}
	}
	return false
}

func (p *Player) CanMine(tileX, tileY int) bool {
	absX := float64(tileX*constants.TILESIZE + constants.TILESIZE/2)
	absY := float64(tileY*constants.TILESIZE + constants.TILESIZE/2)
	return engine.Distance(p.X, p.Y, absX, absY) <= p.Mining.Range
}

func (p *Player) CanBuild(tileX, tileY int) bool {
	absX := float64(tileX*constants.TILESIZE + constants.TILESIZE/2)
	absY := float64(tileY*constants.TILESIZE + constants.TILESIZE/2)
	return engine.Distance(p.X, p.Y, absX, absY) <= p.Building.Range
}

func (p *Player) StartMining(tile *engine.Tile, tileX, tileY int) {
	p.Action = constants.Mining
	p.Mining.TargetTile = &TargetTile{
		ID:    tile.ID,
		TileX: tileX,
		TileY: tileY,
	}
}

func (p *Player) StopMining() {
	p.Action = constants.Idling
	p.Mining.Timer = 0
	p.Mining.TargetTile = nil
}

func (p *Player) StartBuilding(build *engine.Build) {
	p.Action = constants.Building
}

func (p *Player) StopBuilding() {

}

func (p *Player) Draw(screen *ebiten.Image, camera *engine.Camera) {
	relativePositionX := (p.X - float64(p.Sprite.Image.Bounds().Dx()/2)) - (camera.X - float64(screen.Bounds().Dx()/2))
	relativePositionY := (p.Y - float64(p.Sprite.Image.Bounds().Dy()/2)) - (camera.Y - float64(screen.Bounds().Dy()/2))

	p.Sprite.Draw(screen, relativePositionX, relativePositionY, 1.0, 1.0)

	if p.Action == constants.Mining && p.Mining.TargetTile != nil {
		relativePositionX2 := (float64(p.Mining.TargetTile.TileX * constants.TILESIZE)) - (camera.X - float64(screen.Bounds().Dx()/2)) + constants.TILESIZE/2
		relativePositionY2 := (float64(p.Mining.TargetTile.TileY * constants.TILESIZE)) - (camera.Y - float64(screen.Bounds().Dy()/2)) + constants.TILESIZE/2

		vector.StrokeLine(screen,
			float32(relativePositionX+float64(p.Sprite.Image.Bounds().Dx()/2)),
			float32(relativePositionY+float64(p.Sprite.Image.Bounds().Dy()/2)),
			float32(relativePositionX2),
			float32(relativePositionY2),
			3, color.RGBA{0xFF, 0x00, 0x00, 0xFF}, false)
	}
}
