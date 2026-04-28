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

	SpritesInstantiated []*sprite.Sprite
}

func (ui *InventoryUI) Instantiate(scaleX, scaleY float64) {
	for range len(ui.Data.Items) {
		sprite := ui.SpriteSlot.Instantiate()
		
		sprite.ScaleX = scaleX
		sprite.ScaleY = scaleY
		ui.SpritesInstantiated = append(ui.SpritesInstantiated, sprite)
	}
}

func (ui *InventoryUI) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()

	ui.StartX = ( screenWidth - (len(ui.Data.Items)*(ui.Spacing+ui.SpriteSlot.Image.Bounds().Dx())-ui.Spacing) ) / 2
	ui.StartY = screenHeight - ui.SpriteSlot.Image.Bounds().Dy() - 10

	for index, item := range ui.Data.Items {
		sprite := ui.SpritesInstantiated[index]

		sprite.X = ui.StartX + index*(ui.Spacing+ui.SpriteSlot.Image.Bounds().Dx())
		sprite.Y = ui.StartY

		sprite.Draw(screen)

		if item.Item != nil {
			
			itemSprite := item.Item.GetSprite()
			itemSprite.X = sprite.X + (ui.SpriteSlot.Image.Bounds().Dx()-32)/2
			itemSprite.Y = sprite.Y + (ui.SpriteSlot.Image.Bounds().Dy()-32)/2
			itemSprite.ScaleX = 1
			itemSprite.ScaleY = 1

			itemSprite.Draw(screen)
		}
	}
}
