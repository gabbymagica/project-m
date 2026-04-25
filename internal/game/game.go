package game

import (
	"fmt"
	"image/color"
	"project_m/assets"
	inventoryUi "project_m/internal/display/ui"
	sprite "project_m/internal/engine"
	"project_m/internal/game/entities"
	"project_m/internal/game/entities/player"
	item "project_m/internal/game/items"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Object struct {
	Sprite *ebiten.Image
}

type Game struct {
	Map      [32][100]*Object
	Tilesize int

	Player   *player.Player
	Entities []*entities.Entity
	Inputs   []ebiten.Key

	InventoryUi *inventoryUi.InventoryUI
}

func (g *Game) Update() error {
	g.Inputs = inpututil.AppendPressedKeys(g.Inputs[:0])

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.X += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.X -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Y += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Y -= 5
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.DrawTileBorders(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player X: %f\nPlayer Y: %f", g.Player.X, g.Player.Y))

	g.InventoryUi.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1366, 788
}

func GameSetup() *Game {
	ebiten.SetWindowSize(1366, 788)
	ebiten.SetWindowTitle("mindustry 3")

	game := &Game{}
	game.Tilesize = 32


	el_quadrado_vermelho := ebiten.NewImage(16, 16)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	inventory := &inventory.Inventory{}


	game.Player = &player.Player{
		Entity: entities.Entity{
			Sprite: el_quadrado_vermelho,
		},

		Inventory: inventory,
	}

	inventoryUi := &inventoryUi.InventoryUI{
		Data: inventory,
		SpriteSlot: sprite.NewSprite(
			"slot",
			"slot",
			assets.LoadImage("slot.png"),

		),
		StartX : 360,
		StartY : 720,
		SpriteSlotSize : 64,
		Spacing : 10,
	}

	inventoryUi.Instantiate(1, 1)

	game.InventoryUi = inventoryUi

	itemManager := item.NewItemManager()

	cobre := itemManager.Cobre.Instantiate()

	inventory.AddItem(cobre, 5)
	return game
}
