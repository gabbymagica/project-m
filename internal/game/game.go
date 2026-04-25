package game

import (
	"fmt"
	"image/color"
	"project_m/internal/game/entities"
	"project_m/internal/game/entities/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Object struct {
	Sprite *ebiten.Image
}

type Game struct {
	Map      [36][240]*Object
	Tilesize int

	Player   *player.Player
	Entities []*entities.Entity
	Inputs   []ebiten.Key
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
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(960/2, 540/2)

	if g.Player.Sprite != nil {
		screen.DrawImage(g.Player.Sprite, op)
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player X: %f\nPlayer Y: %f", g.Player.X, g.Player.Y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func GameSetup() *Game {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("mindustry 3")

	game := &Game{}
	game.Tilesize = 16

	el_quadrado_vermelho := ebiten.NewImage(16, 16)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	game.Player = &player.Player{
		Entity: entities.Entity{
			Sprite: el_quadrado_vermelho,
		},
	}

	return game
}
