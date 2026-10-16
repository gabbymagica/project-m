package game

import (
	"project_m/assets"
	"project_m/internal/constants"
	"project_m/internal/display/ui"
	inventoryUi "project_m/internal/display/ui"
	"project_m/internal/engine"
	"project_m/internal/game/entities/player"
	"project_m/internal/game/load"
	"project_m/internal/inventory"
)

func (g *Game) InitMap() {
	g.Map = engine.NewMap(constants.TILESIZE, 100, 30)

	grass := engine.NewAutoGenerateTile(load.TILES[constants.TileGrass], 20)
	copper := engine.NewAutoGenerateTile(load.TILES[constants.TileOreCopper], 200)

	prioridades := map[*engine.AutoGenerateTile]float32{
		grass:  500,
		copper: 1,
	}

	g.Map.GenerateMap(prioridades)
}

func (g *Game) InitPlayer() {
	

	inventory := &inventory.Inventory{}

	g.Player = &player.Player{
		Inventory: inventory,
	}
	g.Player.Sprite = load.PLAYER_SPRITES["main"]
	g.Player.Speed = 5
	g.Player.IsAction = false
	g.Player.Mining = &player.Mining{
		IsMining: false,
		TargetTile:   nil,
	}
	g.Player.X = float64(constants.LAYOUT_WIDTH / 2)
	g.Player.Y = float64(constants.LAYOUT_HEIGHT / 2)
	g.Player.Mining.Range = 200
	g.Player.Mining.MiningDelay = 1.0
	g.Player.Mining.MinerableOres = []constants.TileID{constants.TileOreCopper}

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
