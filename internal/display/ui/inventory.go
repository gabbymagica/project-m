package inventoryUi

import (
	sprite "project_m/internal/engine"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
)

type InventoryUI struct {
	Data           *inventory.Inventory
	SpriteSlot     *sprite.Sprite
	StartX         int
	StartY         int
	Spacing        int
	SpriteSlotSize int

	SpritesInstantiated []*sprite.Sprite
}

func (ui *InventoryUI) Instantiate(scaleX, scaleY float64) {
	ui.SpriteSlotSize = int(float64(ui.SpriteSlotSize) * scaleX)
	for index := range ui.Data.Items {
		sprite := ui.SpriteSlot.Instantiate()
		sprite.X = ui.StartX + index*(ui.Spacing+ui.SpriteSlotSize)
		sprite.Y = ui.StartY

		sprite.ScaleX = scaleX
		sprite.ScaleY = scaleY
		ui.SpritesInstantiated = append(ui.SpritesInstantiated, sprite)
	}
}

func (ui *InventoryUI) Draw(screen *ebiten.Image) {
	for index, item := range ui.Data.Items {
		sprite := ui.SpritesInstantiated[index]
		sprite.Draw(screen)

		if item.Item != nil {
			itemSprite := item.Item.GetSprite()
			itemSprite.X = sprite.X + (ui.SpriteSlotSize-32)/2
			itemSprite.Y = sprite.Y + (ui.SpriteSlotSize-32)/2
			itemSprite.ScaleX = 1
			itemSprite.ScaleY = 1

			itemSprite.Draw(screen)

		}
	}
}
