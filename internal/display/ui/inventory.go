package inventoryUi

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
	Item *inventory.InventorySlot
}

type InventoryUI struct {
	Font       *text.GoTextFaceSource
	Data       *inventory.Inventory
	SpriteSlot *sprite.Sprite
	StartX     int
	StartY     int
	Spacing    int

	StateDragDrop StateDragDrop
	Open bool
	SpritesInstantiated []*sprite.Sprite
}

func (ui *InventoryUI) Instantiate(scaleX, scaleY float64) {

	ui.Font = engine.LoadFont("fonts/Pix32.ttf")

	for range len(ui.Data.Items) {
		sprite := ui.SpriteSlot.Instantiate()

		sprite.ScaleX = scaleX
		sprite.ScaleY = scaleY
		ui.SpritesInstantiated = append(ui.SpritesInstantiated, sprite)
	}
}

func (ui *InventoryUI) Drag(mouseX, mouseY int) {
	for index, item := range ui.Data.Items {
		if index == 9 && !ui.Open {
			break
		}
		sprite := ui.SpritesInstantiated[index]

		if mouseX >= sprite.X && mouseX <= sprite.X+sprite.Image.Bounds().Dx() &&
			mouseY >= sprite.Y && mouseY <= sprite.Y+sprite.Image.Bounds().Dy() {
			
			if item.Item == nil {
				break
			}

			ui.StateDragDrop.Active = true
			ui.StateDragDrop.Item = &ui.Data.Items[index]
			break
		}
	}
}

func (ui *InventoryUI) Drop(mouseX, mouseY int) {
	if ui.StateDragDrop.Item == nil {
		return
	}

	for index := range ui.Data.Items {
		if index == 9 && !ui.Open {
			break
		}
		sprite := ui.SpritesInstantiated[index]

		if mouseX >= sprite.X && mouseX <= sprite.X+sprite.Image.Bounds().Dx() &&
			mouseY >= sprite.Y && mouseY <= sprite.Y+sprite.Image.Bounds().Dy() {
			
			itemTrocado := ui.Data.Items[index]
			ui.Data.Items[index] = *ui.StateDragDrop.Item;
			*ui.StateDragDrop.Item = itemTrocado

			break
		}
	}

	ui.StateDragDrop.Active = false
	ui.StateDragDrop.Item = nil
}


func (ui *InventoryUI) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()

	ui.StartX = (screenWidth - (9*(ui.Spacing+ui.SpriteSlot.Image.Bounds().Dx()) - ui.Spacing)) / 2
	ui.StartY = screenHeight - ui.SpriteSlot.Image.Bounds().Dy() - 10

	for index, item := range ui.Data.Items  {
		if index == 9 {
			if !ui.Open {
				break
			} 
			ui.StartY = 200
		} 
 
		sprite := ui.SpritesInstantiated[index]

		sprite.X = ui.StartX + (index % 9)*(ui.Spacing+ui.SpriteSlot.Image.Bounds().Dx())
		sprite.Y = ui.StartY + (index / 9)*(ui.Spacing+ui.SpriteSlot.Image.Bounds().Dy())

		sprite.Draw(screen)

		if item.Item != nil {

			itemSprite := item.Item.GetSprite()
			itemSprite.X = sprite.X + (ui.SpriteSlot.Image.Bounds().Dx()-32)/2
			itemSprite.Y = sprite.Y + (ui.SpriteSlot.Image.Bounds().Dy()-32)/2
			itemSprite.ScaleX = 1
			itemSprite.ScaleY = 1

			itemSprite.Draw(screen)

			engine.DrawText(screen, ui.Font, fmt.Sprintf("%d", item.Quantity), float64(sprite.X+10), float64(sprite.Y+6), 10, color.White)

		}
		
	}

	if ui.StateDragDrop.Active && ui.StateDragDrop.Item != nil {
		mouseX, mouseY := ebiten.CursorPosition()
		itemSprite := ui.StateDragDrop.Item.Item.GetSprite()
		itemSprite.X = mouseX - 16
		itemSprite.Y = mouseY - 16
		itemSprite.Draw(screen)
	}
}
