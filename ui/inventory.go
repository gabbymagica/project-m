package inventoryUi

import (
	sprite "project_m/engine"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
)

type InventoryUI struct {
	Data *inventory.Inventory
	SpriteSlot *sprite.Sprite
	StartX int
	StartY int
	Spacing int
	SpriteSlotSize int

	SpritesInstantiated []*sprite.Sprite
}

func (ui *InventoryUI) Instantiate() {
	for index := range ui.Data.Items {
		sprite := ui.SpriteSlot.Instantiate()
		sprite.X = ui.StartX + index*(ui.Spacing+ui.SpriteSlotSize)
		sprite.Y = ui.StartY

		ui.SpritesInstantiated = append(ui.SpritesInstantiated, sprite)
	}
}

func (ui *InventoryUI) Draw(screen *ebiten.Image) {
	for _, sprite := range ui.SpritesInstantiated {
		sprite.Draw(screen)
	}
}

func (ui *InventoryUI) Scale(scaleX, scaleY float64) {
	ui.SpriteSlotSize = int(float64(ui.SpriteSlotSize) * scaleX)
	for _, sprite := range ui.SpritesInstantiated {
		sprite.Scale(scaleX, scaleY)
	}
}
	