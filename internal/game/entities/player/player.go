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

type Destroying struct {
	DestroyTarget *engine.Build
	Range         float64
}

type Player struct {
	engine.Entity
	Inventory  *inventory.Inventory
	Action     constants.ActionID
	Mining     Mining
	Building   Building
	Destroying Destroying
}

func (p *Player) Update(dt float64) {
	p.UpdateMining(dt)
	p.UpdateBuilding(dt)
	p.UpdateDestroying(dt)
}

func (p *Player) StopAllActions() {
	p.StopMining()
	p.StopBuilding()
	p.StopDestroying()
	p.Action = constants.Idling
}

func (p *Player) TogglePreview(previewAction constants.ActionID) {
	if p.Action == previewAction {
		p.StopAllActions()
		return
	}

	p.StopAllActions()
	p.Action = previewAction
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

func (p *Player) UpdateBuilding(dt float64) {
	var buildTarget = p.Building.BuildTarget
	if buildTarget == nil {
		return
	}
	if p.Action == constants.Building && buildTarget.FinishedBuilding == false {
		if p.CanBuild(buildTarget.TileX, buildTarget.TileY, buildTarget.BuildType) == false {
			p.StopBuilding()
			return
		}

		buildTarget.BuildProgress += dt
		if buildTarget.BuildProgress >= buildTarget.BuildType.BuildingTimeSeconds {
			buildTarget.FinishedBuilding = true
			p.StopBuilding()
		}
	}
}

func (p *Player) UpdateDestroying(dt float64) {
	var destroyTarget = p.Destroying.DestroyTarget
	if destroyTarget == nil {
		return
	}
	if p.Action == constants.Destroying {
		if p.CanDestroy(destroyTarget) == false {
			p.StopDestroying()
			return
		}

		destroyTarget.BuildProgress -= dt * 2
		if destroyTarget.BuildProgress <= 0 {
			destroyTarget.Destroyed = true
			p.StopDestroying()
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

func (p *Player) CanBuild(tileX, tileY int, buildType *engine.BuildType) bool {
	if buildType == nil {
		return false
	}

	objectX, objectY := buildType.CenterPosition(tileX, tileY, constants.TILESIZE)
	return engine.Distance(p.X, p.Y, objectX, objectY) <= p.Building.Range
}

func (p *Player) CanDestroy(build *engine.Build) bool {
	if build == nil || build.BuildType == nil || build.Destroyed {
		return false
	}

	objectX, objectY := build.CenterPosition(constants.TILESIZE)
	return engine.Distance(p.X, p.Y, objectX, objectY) <= p.Destroying.Range
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
	p.Building.BuildTarget = build
}

func (p *Player) StopBuilding() {
	p.Action = constants.Idling
	p.Building.BuildTarget = nil
}

func (p *Player) StartDestroying(build *engine.Build) {
	p.Action = constants.Destroying
	build.FinishedBuilding = false
	p.Destroying.DestroyTarget = build
}

func (p *Player) StopDestroying() {
	p.Action = constants.Idling
	p.Destroying.DestroyTarget = nil
}

func (p *Player) drawLaser(screen *ebiten.Image, camera *engine.Camera, targetWorldX, targetWorldY float64, clr color.Color) {
	startX, startY := camera.WorldToScreen(p.X, p.Y)
	endX, endY := camera.WorldToScreen(targetWorldX, targetWorldY)

	vector.StrokeLine(
		screen,
		float32(startX), float32(startY),
		float32(endX), float32(endY),
		3, clr, false,
	)
}

func (p *Player) Draw(screen *ebiten.Image, camera *engine.Camera) {
	halfW := float64(p.Sprite.Image.Bounds().Dx()) / 2
	halfH := float64(p.Sprite.Image.Bounds().Dy()) / 2
	drawX, drawY := camera.WorldToScreen(p.X-halfW, p.Y-halfH)

	p.Sprite.Draw(screen, drawX, drawY, 1.0, 1.0, 1.0)

	switch p.Action {
	case constants.Mining:
		if p.Mining.TargetTile != nil {
			targetX := float64(p.Mining.TargetTile.TileX*constants.TILESIZE + constants.TILESIZE/2)
			targetY := float64(p.Mining.TargetTile.TileY*constants.TILESIZE + constants.TILESIZE/2)
			p.drawLaser(screen, camera, targetX, targetY, color.RGBA{0xFF, 0x00, 0x00, 0xFF})
		}
	case constants.Building:
		if p.Building.BuildTarget != nil {
			targetX, targetY := p.Building.BuildTarget.CenterPosition(constants.TILESIZE)
			p.drawLaser(screen, camera, targetX, targetY, color.RGBA{0x00, 0xFF, 0x00, 0xFF})
		}
	case constants.Destroying:
		if p.Destroying.DestroyTarget != nil {
			targetX, targetY := p.Destroying.DestroyTarget.CenterPosition(constants.TILESIZE)
			p.drawLaser(screen, camera, targetX, targetY, color.RGBA{0xFF, 0x00, 0x00, 0xFF})
		}
	}
}
