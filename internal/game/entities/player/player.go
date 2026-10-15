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

type Mining struct {
	IsMining    bool
	Object      *engine.Object
	Range       float64
	MiningDelay float64
	Timer       float64
}

type Building struct {
	IsBuilding    bool
	Object        *engine.Object
	Range         float64
	BuildingDelay float64
	Timer         float64
}

type Player struct {
	engine.Entity
	Inventory *inventory.Inventory
	IsAction  bool
	Mining    *Mining
}

func (p *Player) Update(dt float64) {
	p.UpdateMining(dt)
}

func (p *Player) UpdateMining(dt float64) {
	if p.IsAction && p.Mining.IsMining && p.Mining.Object != nil {
		objectX := float64(p.Mining.Object.TileX*constants.TILESIZE + constants.TILESIZE/2)
		objectY := float64(p.Mining.Object.TileY*constants.TILESIZE + constants.TILESIZE/2)
		if engine.Distance(p.X, p.Y, objectX, objectY) > p.Mining.Range {
			p.IsAction = false
			p.Mining.IsMining = false
			p.Mining.Object = nil
			return
		}

		p.Mining.Timer += dt
		if p.Mining.Timer >= p.Mining.MiningDelay {
			p.Mining.Timer = 0
			p.Inventory.AddItem(item.Manager.Cobre, 1)
		}

	}
}

func (p *Player) StartMining(object *engine.Object) {
	objectX := float64(object.TileX*constants.TILESIZE + constants.TILESIZE/2)
	objectY := float64(object.TileY*constants.TILESIZE + constants.TILESIZE/2)
	if engine.Distance(p.X, p.Y, objectX, objectY) <= p.Mining.Range {
		p.Mining.IsMining = true
		p.IsAction = true
		p.Mining.Object = object
		p.Mining.Timer = 0
	}
}

func (p *Player) Draw(screen *ebiten.Image, camera *engine.Camera) {
	relativePositionX := (p.X - float64(p.Sprite.Image.Bounds().Dx()/2)) - (camera.X - float64(screen.Bounds().Dx()/2))
	relativePositionY := (p.Y - float64(p.Sprite.Image.Bounds().Dy()/2)) - (camera.Y - float64(screen.Bounds().Dy()/2))

	p.Sprite.Draw(screen, relativePositionX, relativePositionY, 1.0, 1.0)

	if p.IsAction && p.Mining.IsMining && p.Mining.Object != nil {
		relativePositionX2 := (float64(p.Mining.Object.TileX * constants.TILESIZE)) - (camera.X - float64(screen.Bounds().Dx()/2)) + constants.TILESIZE/2
		relativePositionY2 := (float64(p.Mining.Object.TileY * constants.TILESIZE)) - (camera.Y - float64(screen.Bounds().Dy()/2)) + constants.TILESIZE/2

		vector.StrokeLine(screen,
			float32(relativePositionX+float64(p.Sprite.Image.Bounds().Dx()/2)),
			float32(relativePositionY+float64(p.Sprite.Image.Bounds().Dy()/2)),
			float32(relativePositionX2),
			float32(relativePositionY2),
			3, color.RGBA{0xFF, 0x00, 0x00, 0xFF}, false)
	}
}
