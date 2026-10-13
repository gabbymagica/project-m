package game

import (
	"image/color"
	"time"

	"project_m/internal/constants"
	inventoryUi "project_m/internal/display/ui"
	"project_m/internal/engine"
	"project_m/internal/game/entities/player"
	item "project_m/internal/game/items"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Map *engine.Map

	Camera   *engine.Camera
	Player   *player.Player
	InventoryUi *inventoryUi.InventoryUI
	Entities []*engine.Entity
	Inputs   []ebiten.Key

	TestObject *engine.Object

	GameMode    GameMode
	RenderFlags *RenderFlags

	ItemManager *item.ItemManager


	// Time
	lastUpdate time.Time
	deltaTime float64
	// Time
}

func (g *Game) CalculateDeltaTime() {
	currentTime := time.Now()
	if !g.lastUpdate.IsZero() {
		g.deltaTime = currentTime.Sub(g.lastUpdate).Seconds()
	} else {
		g.deltaTime = 0
	}
	g.lastUpdate = currentTime

	// (Opcional de segurança: trava o delta time máximo caso o jogo sofra um "lag spike" ou minimize)
	if g.deltaTime > 0.1 {
        g.deltaTime = 0.1
    }
}

func (g *Game) Update() error {
	g.CalculateDeltaTime()

	g.HandleGameInputs()
	g.Player.Update(g.deltaTime)
	g.Camera.Follow(g.Player.X, g.Player.Y)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.HandleTileDrawStates(screen)
	g.Player.Draw(screen, g.Camera)
	g.InventoryUi.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.LAYOUT_WIDTH, constants.LAYOUT_HEIGHT
}



func GameSetup() *Game {
	ebiten.SetWindowSize(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("mindustry 3")

	game := &Game{}

	game.RenderFlags = &RenderFlags{
		ShowSprites: true,
	}

	game.Camera = &engine.Camera{
		ScreenWidth:  constants.LAYOUT_WIDTH,
		ScreenHeight: constants.LAYOUT_HEIGHT,
	}
	
	game.InitMap()
	game.InitPlayer()
	game.InitUi()

	// Teste
	game.Player.Inventory.AddItem(item.Manager.Cobre, 10)
	// Teste

	el_quadrado_vermelho := ebiten.NewImage(16, 16)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	objeto := engine.NewObject("quadradovermelho", 0, 0, 0, 3, 3, 0, engine.NewSprite("quadradovermelho", "quadradovermelho", el_quadrado_vermelho))

	game.Map.PutObject(objeto)

	game.TestObject = objeto.Clone()


	return game
}
