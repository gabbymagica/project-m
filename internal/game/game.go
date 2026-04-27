package game

import (
	"fmt"
	"image/color"
	"project_m/assets"
	inventoryUi "project_m/internal/display/ui"
	"project_m/internal/engine"
	"project_m/internal/game/entities/player"
	item "project_m/internal/game/items"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	Map *engine.Map

	Player   *player.Player
	Entities []*engine.Entity
	Inputs   []ebiten.Key

	InventoryUi *inventoryUi.InventoryUI
}

func (g *Game) Update() error {
	g.Inputs = inpututil.AppendPressedKeys(g.Inputs[:0])

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.X += 30
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.X -= 30
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Y += 30
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Y -= 30
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawTileBorders(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player X: %f\nPlayer Y: %f", g.Player.X, g.Player.Y))

	g.InventoryUi.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1366, 768
}

func GameSetup() *Game {
	ebiten.SetWindowSize(1366, 768)
	ebiten.SetWindowTitle("mindustry 3")

	game := &Game{}
	game.Map = engine.NewMap(16)

	el_quadrado_vermelho := ebiten.NewImage(16, 16)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	inventory := &inventory.Inventory{}

	game.Player = &player.Player{
		Inventory: inventory,
	}

	inventoryUi := &inventoryUi.InventoryUI{
		Data: inventory,
		SpriteSlot: engine.NewSprite(
			"slot",
			"slot",
			assets.LoadImage("slot.png"),
		),
		StartX:         360,
		StartY:         720,
		SpriteSlotSize: 64,
		Spacing:        10,
	}

	inventoryUi.Instantiate(1, 1)

	game.InventoryUi = inventoryUi

	itemManager := item.NewItemManager()

	cobre := itemManager.Cobre.Instantiate()

	inventory.AddItem(cobre, 5)

	game.Map.NewObject(0, 0, 1, 1, engine.NewSprite("", "", el_quadrado_vermelho))

	return game
}
