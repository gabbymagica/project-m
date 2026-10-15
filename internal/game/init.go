package game

import (
	"image/color"
	"project_m/assets"
	"project_m/internal/constants"
	"project_m/internal/display/ui"
	inventoryUi "project_m/internal/display/ui"
	"project_m/internal/engine"
	"project_m/internal/game/entities/player"
	"project_m/internal/inventory"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) InitMap() {
	g.Map = engine.NewMap(constants.TILESIZE)

	grass := engine.NewObject(constants.TileGrass, constants.LAYER_GROUND, 1, 1, 10, engine.NewSprite("grass", "grass", assets.LoadImage("tiles/grass.png")))
	copper := engine.NewObject(constants.OreCopper, constants.LAYER_GROUND, 1, 1, 200, engine.NewSprite("copper_tile", "copper_tile", assets.LoadImage("tiles/copper.png")))

	prioridades := map[*engine.Object]float32{
		grass:  500,
		copper: 1,
	}

	g.Map.GenerateMap(prioridades)
}

func (g *Game) InitPlayer() {
	el_quadrado_preto := ebiten.NewImage(32, 32)
	el_quadrado_preto.Fill(color.RGBA{0, 0, 0, 0xff})
	spritePlayer := engine.NewSprite("player", "player", el_quadrado_preto)

	inventory := &inventory.Inventory{}

	g.Player = &player.Player{
		Inventory: inventory,
	}
	g.Player.Sprite = spritePlayer
	g.Player.Speed = 5
	g.Player.IsAction = false
	g.Player.Mining = &player.Mining{
		IsMining: false,
		Object:   nil,
	}
	g.Player.X = float64(constants.LAYOUT_WIDTH / 2)
	g.Player.Y = float64(constants.LAYOUT_HEIGHT / 2)
	g.Player.Mining.Range = 200
	g.Player.Mining.MiningDelay = 1.0
	g.Player.Mining.MinerableOres = []constants.GroundID{constants.OreCopper}

}

func (g *Game) InitUi() {
	g.GameFonts = ui.InitFonts()

	inventoryUi := &inventoryUi.InventoryUI{
		Font: g.GameFonts.InventoryUIFont,
		Data: g.Player.Inventory,
		SpriteSlot: engine.NewSprite(
			"slot",
			"slot",
			assets.LoadImage("slot.png"),
		),
		Spacing: 10,
	}
	inventoryUi.Instantiate()
	g.InventoryUi = inventoryUi

}
