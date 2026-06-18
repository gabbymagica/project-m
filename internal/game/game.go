package game

import (
	"image/color"
	"project_m/assets"
	inventoryUi "project_m/internal/display/ui"
	"project_m/internal/engine"
	"project_m/internal/game/entities/player"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map *engine.Map

	Camera   *engine.Camera
	Player   *player.Player
	Entities []*engine.Entity
	Inputs   []ebiten.Key

	TestObject *engine.Object

	GameMode    GameMode
	RenderFlags *RenderFlags
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

	game.RenderFlags = &RenderFlags{
		ShowSprites: true,
	}

	game.Map = engine.NewMap(32)

	grass := engine.NewObject("grass", 0, 0, 0, 1, 1, 10, engine.NewSprite("grass", "grass", assets.LoadImage("tiles/grass.png")))
	rock_grass := engine.NewObject("rock_grass", 0, 0, 0, 1, 1, 10, engine.NewSprite("rock_grass", "rock_grass", assets.LoadImage("tiles/rock-grass.png")))
	copper := engine.NewObject("copper", 0, 0, 0, 1, 1, 400, engine.NewSprite("copper_tile", "copper_tile", assets.LoadImage("tiles/copper.png")))

	prioridades := map[*engine.Object]float32{
		grass:      500,
		rock_grass: 100,
		copper:     1,
	}

	game.Map.GenerateMap(prioridades)

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
		Spacing: 10,
	}
	inventoryUi.Instantiate(1, 1)
	game.InventoryUi = inventoryUi

	el_quadrado_vermelho := ebiten.NewImage(16, 16)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	objeto := engine.NewObject("quadradovermelho", 0, 0, 0, 3, 3, 0, engine.NewSprite("quadradovermelho", "quadradovermelho", el_quadrado_vermelho))

	game.Map.PutObject(objeto)

	game.TestObject = objeto.Clone()

	return game
}
