package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Entity struct {
	X, Y   float64
	Sprite *ebiten.Image
}

type Game struct {
	Player   *Entity
	Entities []*Entity
	Inputs   []ebiten.Key
}

func (g *Game) Update() error {
	g.Inputs = inpututil.AppendPressedKeys(g.Inputs[:0])

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.Player.X += 2.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.Player.X -= 2.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.Player.Y += 2.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.Player.Y -= 2.5
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.Player.X, g.Player.Y)
	if g.Player.Sprite != nil {
		screen.DrawImage(g.Player.Sprite, op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Player X: %f\nPlayer Y: %f", g.Player.X, g.Player.Y))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 960, 540
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("mindustry 3")

	game := &Game{}

	el_quadrado_vermelho := ebiten.NewImage(32, 32)
	el_quadrado_vermelho.Fill(color.RGBA{0xff, 0, 0, 0xff})

	game.Player = &Entity{
		Sprite: el_quadrado_vermelho,
		X:      10,
		Y:      10,
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
