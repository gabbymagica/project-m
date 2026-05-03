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
)

type TileStateInt int

const (
	StateDestroying TileStateInt = iota
	StateBuilding
	StatePreview
	StateDrawSprites
)

type Game struct {
	Map *engine.Map

	Camera   *engine.Camera
	Player   *player.Player
	Entities []*engine.Entity
	Inputs   []ebiten.Key

	TestObject *engine.Object

	TileStates  map[TileStateInt]bool
	InventoryUi *inventoryUi.InventoryUI
}

const MINIMUM_TILESIZE_SIZE = 1

func (g *Game) Update() error {
	g.HandleGameInputs()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.HandleTileDrawStates(screen)
	g.InventoryUi.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1366, 768
}

func GameSetup() *Game {
	ebiten.SetWindowSize(1366, 768)
	ebiten.SetWindowTitle("mindustry 3")

	game := &Game{}
	game.TileStates = map[TileStateInt]bool{
		StateBuilding:    false,
		StateDestroying:  false,
		StatePreview:     false,
		StateDrawSprites: true,
	}

	game.Map = engine.NewMap(32)

	el_quadrado_vermelho := ebiten.NewImage(16, 16)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	inventory := &inventory.Inventory{}

	game.Player = &player.Player{
		Inventory: inventory,
	}

	game.Camera = &engine.Camera{}

	inventoryUi := &inventoryUi.InventoryUI{
		Data: inventory,
		SpriteSlot: engine.NewSprite(
			"slot",
			"slot",
			assets.LoadImage("slot.png"),
		),
		StartX:  360,
		StartY:  720,
		Spacing: 10,
	}

	inventoryUi.Instantiate(1, 1)

	game.InventoryUi = inventoryUi

	itemManager := item.NewItemManager()

	cobre := itemManager.Cobre.Instantiate()

	inventory.AddItem(cobre, 5)

	objeto, err := game.Map.NewObject(0, 0, 0, 3, 3, engine.NewSprite("", "", el_quadrado_vermelho))
	clone := *objeto
	clone.Sprite = objeto.Sprite.Instantiate()

	game.TestObject = &clone
	fmt.Printf("%v", err)

	return game
}
