package ui

import (
	"fmt"
	"image/color"
	"project_m/internal/engine"
	sprite "project_m/internal/engine"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type StateDragDrop struct {
	Active bool
	Item   *inventory.InventorySlot
}

type InventoryUI struct {
	Font       *text.GoTextFaceSource
	Data       *inventory.Inventory
	SpriteSlot *sprite.Sprite
	StartX     float64
	StartY     float64
	Spacing    int

	StateDragDrop       StateDragDrop
	Open                bool
	SpritesInstantiated []*sprite.Sprite
}

func (ui *InventoryUI) Instantiate() {
	for range len(ui.Data.Items) {
		sprite := ui.SpriteSlot.Instantiate()

		ui.SpritesInstantiated = append(ui.SpritesInstantiated, sprite)
	}
}

// func (ui *InventoryUI) Drag(mouseX, mouseY int) {
// 	for index, item := range ui.Data.Items {
// 		if index == 9 && !ui.Open {
// 			break
// 		}
// 		sprite := ui.SpritesInstantiated[index]

// 		if mouseX >= sprite.X && mouseX <= sprite.X+sprite.Image.Bounds().Dx() &&
// 			mouseY >= sprite.Y && mouseY <= sprite.Y+sprite.Image.Bounds().Dy() {

// 			if item.Item == nil {
// 				break
// 			}

// 			ui.StateDragDrop.Active = true
// 			ui.StateDragDrop.Item = &ui.Data.Items[index]
// 			break
// 		}
// 	}
// }

// func (ui *InventoryUI) Drop(mouseX, mouseY int) {
// 	if ui.StateDragDrop.Item == nil {
// 		return
// 	}

// 	for index := range ui.Data.Items {
// 		if index == 9 && !ui.Open {
// 			break
// 		}
// 		sprite := ui.SpritesInstantiated[index]

// 		if mouseX >= sprite.X && mouseX <= sprite.X+sprite.Image.Bounds().Dx() &&
// 			mouseY >= sprite.Y && mouseY <= sprite.Y+sprite.Image.Bounds().Dy() {

// 			itemTrocado := ui.Data.Items[index]
// 			ui.Data.Items[index] = *ui.StateDragDrop.Item;
// 			*ui.StateDragDrop.Item = itemTrocado

// 			break
// 		}
// 	}

// 	ui.StateDragDrop.Active = false
// 	ui.StateDragDrop.Item = nil
// }

func (ui *InventoryUI) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()

	ui.StartX = float64(screenWidth-(9*(ui.Spacing+ui.SpriteSlot.Image.Bounds().Dx())-ui.Spacing)) / 2
	ui.StartY = float64(screenHeight) - float64(ui.SpriteSlot.Image.Bounds().Dy()) - 10

	for index, item := range ui.Data.Items {
		if index == 9 {
			if !ui.Open {
				break
			}
			ui.StartY = 200
		}

		sprite := ui.SpritesInstantiated[index]

		spriteX := ui.StartX + float64(index%9)*(float64(ui.Spacing)+float64(ui.SpriteSlot.Image.Bounds().Dx()))
		spriteY := ui.StartY + float64(index/9)*(float64(ui.Spacing)+float64(ui.SpriteSlot.Image.Bounds().Dy()))

		sprite.Draw(screen, spriteX, spriteY, 1.0, 1.0)

		if item.Item != nil {

			itemSprite := item.Item.GetSprite()
			itemSpriteX := spriteX + float64(ui.SpriteSlot.Image.Bounds().Dx()-32)/2
			itemSpriteY := spriteY + float64(ui.SpriteSlot.Image.Bounds().Dy()-32)/2

			itemSprite.Draw(screen, itemSpriteX, itemSpriteY, 1.0, 1.0)

			engine.DrawText(screen, ui.Font, fmt.Sprintf("%d", item.Quantity), float64(int(spriteX+10)), float64(int(spriteY+6)), 12, color.White)

		}

	}

	// if ui.StateDragDrop.Active && ui.StateDragDrop.Item != nil {
	// 	mouseX, mouseY := ebiten.CursorPosition()
	// 	itemSprite := ui.StateDragDrop.Item.Item.GetSprite()
	// 	itemSprite.X = mouseX - 16
	// 	itemSprite.Y = mouseY - 16
	// 	itemSprite.Draw(screen)
	// }
}
